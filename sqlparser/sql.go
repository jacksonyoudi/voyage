package sqlparser

import (
	"fmt"
	"github.com/jacksonyoudi/voyage/sqlparser/query"
	"regexp"
	"strings"
)

type step int

const (
	stepType step = iota
	stepSelectField
	stepSelectFrom
	stepSelectComma
	stepSelectFromTable
	stepInsertTable
	stepInsertFieldsOpeningParens
	stepInsertFields
	stepInsertFieldsCommaOrClosingParens
	stepInsertValuesOpeningParens
	stepInsertValuesRWord
	stepInsertValues
	stepInsertValuesCommaOrClosingParens
	stepInsertValuesCommaBeforeOpeningParens
	stepUpdateTable
	stepUpdateSet
	stepUpdateField
	stepUpdateEquals
	stepUpdateValue
	stepUpdateComma
	stepDeleteFromTable
	stepWhere
	stepWhereField
	stepWhereOperator
	stepWhereValue
	stepWhereAnd
)

type parser struct {
	i               int          // 当前所在查询字符串中的位置
	sql             string       // 原sql
	step            step         // 状态机的步骤
	query           *query.Query // 解析后的结构体
	err             error
	nextUpdateField string
}

var reservedWords = []string{
	"(", ")", ">=", "<=", "!=", ",", "=", ">", "<", "SELECT", "INSERT INTO", "VALUES", "UPDATE", "DELETE FROM",
	"WHERE", "FROM", "SET", "AS",
}

// 多余
func Parse(sqls string) (*query.Query, error) {
	qs, err := ParseMany([]string{sqls})
	if len(qs) == 0 {
		return &query.Query{}, nil
	}
	return qs[0], err
}

func ParseMany(sqls []string) ([]*query.Query, error) {
	qs := []*query.Query{}
	for _, sql := range sqls {
		q, err := parse(sql)
		if err != nil {
			return nil, err
		}
		qs = append(qs, q)
	}
	return qs, nil
}

func parse(sql string) (*query.Query, error) {
	return (&parser{
		i:               0,
		sql:             sql,
		step:            1,
		query:           &query.Query{},
		err:             nil,
		nextUpdateField: "",
	}).parse()

}

func (p *parser) parse() (*query.Query, error) {
	q, err := p.doParse()
	p.err = err
	if p.err == nil {
		p.err = p.validate()
	}
	p.logError()
	return q, p.err
}

// 有限状态机
func (p *parser) doParse() (*query.Query, error) {
	for {
		if p.i >= len(p.sql) {
			return p.query, p.err
		}
		switch p.step {
		case stepType:
			switch strings.ToUpper(p.peek()) {
			case "SELECT":
				p.query.Type = query.Select
				p.pop()
				// 下一步
				p.step = stepSelectField
			case "INSERT INTO":
				p.query.Type = query.Insert
				p.pop()
				p.step = stepInsertTable
			case "UPDATE":
				p.query.Type = query.Update
				p.query.Updates = map[string]string{}
				p.pop()
				p.step = stepUpdateTable
			case "DELETE FROM":
				p.query.Type = query.Delete
				p.pop()
				p.step = stepDeleteFromTable
			default:
				return p.query, fmt.Errorf("invalid query type")
			}
			// select语句
		case stepSelectField:
			identifier := p.peek()
			if !isIdentifierOrAsterisk(identifier) {
				return p.query, fmt.Errorf("at SELECT: expected field to SELECT")
			}
			p.query.Fields = append(p.query.Fields, identifier)
			p.pop()
			maybeFrom := p.peek()
			if strings.ToUpper(maybeFrom) == "AS" {
				p.pop()
				alias := p.peek()
				if !isIdentifier(alias) {
					return p.query, fmt.Errorf("at SELECT: expected field alias for \"" + identifier + " as\" to SELECT")
				}
				if p.query.Aliases == nil {
					p.query.Aliases = make(map[string]string)
				}
				// 和原字段关联上
				p.query.Aliases[identifier] = alias
				p.pop()
				maybeFrom = p.peek()
			}
			if strings.ToUpper(maybeFrom) == "FROM" {
				p.step = stepSelectFrom
				continue
			}
			p.step = stepSelectComma

		case stepSelectComma:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.query, fmt.Errorf("at SELECT: expected comma or FROM")
			}
			p.pop()
			p.step = stepSelectField

		case stepSelectFromTable:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.query, fmt.Errorf("at SELECT: expected quoted table name")
			}
			p.query.TableName = tableName
			p.pop()
			p.step = stepWhere

		case stepInsertTable:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.query, fmt.Errorf("at INSERT INTO: expected quoted table name")
			}
			p.query.TableName = tableName
			p.pop()
			p.step = stepInsertFieldsOpeningParens

		case stepDeleteFromTable:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.query, fmt.Errorf("at DELETE FROM: expected quoted table name")
			}
			p.query.TableName = tableName
			p.pop()
			p.step = stepWhere
		case stepUpdateTable:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.query, fmt.Errorf("at UPDATE: expected quoted table name")
			}
			p.query.TableName = tableName
			p.pop()
			p.step = stepUpdateSet

		case stepUpdateSet:
			setRWord := p.peek()
			if setRWord != "SET" {
				return p.query, fmt.Errorf("at UPDATE: expected 'SET'")
			}
			p.pop()
			p.step = stepUpdateField
		case stepUpdateField:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.query, fmt.Errorf("at UPDATE: expected at least one field to update")
			}
			p.nextUpdateField = identifier
			p.pop()
			p.step = stepUpdateEquals

		case stepUpdateEquals:
			equalsRWord := p.peek()
			if equalsRWord != "=" {
				return p.query, fmt.Errorf("at UPDATE: expected '='")
			}
			p.pop()
			p.step = stepUpdateValue

		case stepUpdateValue:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.query, fmt.Errorf("at UPDATE: expected quoted value")
			}
			p.query.Updates[p.nextUpdateField] = quotedValue
			p.nextUpdateField = ""
			p.pop()
			maybeWhere := p.peek()
			if strings.ToUpper(maybeWhere) == "WHERE" {
				p.step = stepWhere
				continue
			}
			p.step = stepUpdateComma

		case stepUpdateComma:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.query, fmt.Errorf("at UPDATE: expected ','")
			}
			p.pop()
			p.step = stepUpdateField

		case stepWhere:
			whereRWord := p.peek()
			if strings.ToUpper(whereRWord) != "WHERE" {
				return p.query, fmt.Errorf("expected WHERE")
			}
			p.pop()
			p.step = stepWhereField
		case stepWhereField:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.query, fmt.Errorf("at WHERE: expected field")
			}
			p.query.Conditions = append(p.query.Conditions, query.Condition{Operand1: identifier, Operand1IsField: true})
			p.pop()
			p.step = stepWhereOperator
		case stepWhereOperator:
			operator := p.peek()
			currentCondition := p.query.Conditions[len(p.query.Conditions)-1]
			switch operator {
			case "=":
				currentCondition.Operator = query.Eq
			case ">":
				currentCondition.Operator = query.Gt
			case ">=":
				currentCondition.Operator = query.Gte
			case "<":
				currentCondition.Operator = query.Lt
			case "<=":
				currentCondition.Operator = query.Lte
			case "!=":
				currentCondition.Operator = query.Ne
			default:
				return p.query, fmt.Errorf("at WHERE: unknown operator")
			}
			p.query.Conditions[len(p.query.Conditions)-1] = currentCondition
			p.pop()
			p.step = stepWhereValue
		case stepWhereValue:
			currentCondition := p.query.Conditions[len(p.query.Conditions)-1]
			identifier := p.peek()
			if isIdentifier(identifier) {
				currentCondition.Operand2 = identifier
				currentCondition.Operand2IsField = true
			} else {
				quotedValue, ln := p.peekQuotedStringWithLength()
				if ln == 0 {
					return p.query, fmt.Errorf("at WHERE: expected quoted value")
				}
				currentCondition.Operand2 = quotedValue
				currentCondition.Operand2IsField = false
			}
			p.query.Conditions[len(p.query.Conditions)-1] = currentCondition
			p.pop()
			p.step = stepWhereAnd
		case stepWhereAnd:
			andRWord := p.peek()
			if strings.ToUpper(andRWord) != "AND" {
				return p.query, fmt.Errorf("expected AND")
			}
			p.pop()
			p.step = stepWhereField
		case stepInsertFieldsOpeningParens:
			openingParens := p.peek()
			if len(openingParens) != 1 || openingParens != "(" {
				return p.query, fmt.Errorf("at INSERT INTO: expected opening parens")
			}
			p.pop()
			p.step = stepInsertFields
		case stepInsertFields:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.query, fmt.Errorf("at INSERT INTO: expected at least one field to insert")
			}
			p.query.Fields = append(p.query.Fields, identifier)
			p.pop()
			p.step = stepInsertFieldsCommaOrClosingParens
		case stepInsertFieldsCommaOrClosingParens:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.query, fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}
			p.pop()
			if commaOrClosingParens == "," {
				p.step = stepInsertFields
				continue
			}
			p.step = stepInsertValuesRWord
		case stepInsertValuesRWord:
			valuesRWord := p.peek()
			if strings.ToUpper(valuesRWord) != "VALUES" {
				return p.query, fmt.Errorf("at INSERT INTO: expected 'VALUES'")
			}
			p.pop()
			p.step = stepInsertValuesOpeningParens
		case stepInsertValuesOpeningParens:
			openingParens := p.peek()
			if openingParens != "(" {
				return p.query, fmt.Errorf("at INSERT INTO: expected opening parens")
			}
			p.query.Inserts = append(p.query.Inserts, []string{})
			p.pop()
			p.step = stepInsertValues
		case stepInsertValues:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.query, fmt.Errorf("at INSERT INTO: expected quoted value")
			}
			p.query.Inserts[len(p.query.Inserts)-1] = append(p.query.Inserts[len(p.query.Inserts)-1], quotedValue)
			p.pop()
			p.step = stepInsertValuesCommaOrClosingParens
		case stepInsertValuesCommaOrClosingParens:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.query, fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}
			p.pop()
			if commaOrClosingParens == "," {
				p.step = stepInsertValues
				continue
			}
			currentInsertRow := p.query.Inserts[len(p.query.Inserts)-1]
			if len(currentInsertRow) < len(p.query.Fields) {
				return p.query, fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}
			p.step = stepInsertValuesCommaBeforeOpeningParens
		case stepInsertValuesCommaBeforeOpeningParens:
			commaRWord := p.peek()
			if strings.ToUpper(commaRWord) != "," {
				return p.query, fmt.Errorf("at INSERT INTO: expected comma")
			}
			p.pop()
			p.step = stepInsertValuesOpeningParens
		}

	}
}

//
func (p *parser) peek() string {
	peeked, _ := p.peekWithLength()
	return peeked
}

func (p *parser) pop() string {
	peeked, l := p.peekWithLength()
	p.i += l
	p.popWhitespace()
	return peeked
}

func (p *parser) popWhitespace() {
	for ; p.i < len(p.sql) && p.sql[p.i] == ' '; p.i++ {
	}
}

func (p *parser) peekWithLength() (string, int) {
	if p.i >= len(p.sql) {
		return "", 0
	}

	// 切割关键词
	for _, rWord := range reservedWords {
		token := strings.ToUpper(p.sql[p.i:min(len(p.sql), p.i+len(rWord))])
		if token == rWord {
			return token, len(token)
		}
	}
	if p.sql[p.i] == '\'' { // Quoted string
		return p.peekQuotedStringWithLength()
	}
	return p.peekIdentifierWithLength()
}

func (p *parser) peekQuotedStringWithLength() (string, int) {
	if len(p.sql) < p.i || p.sql[p.i] != '\'' {
		return "", 0
	}
	for i := p.i + 1; i < len(p.sql); i++ {
		if p.sql[i] == '\'' && p.sql[i-1] != '\\' {
			return p.sql[p.i+1 : i], len(p.sql[p.i+1:i]) + 2 // +2 for the two quotes
		}
	}
	return "", 0
}
func (p *parser) peekIdentifierWithLength() (string, int) {
	for i := p.i; i < len(p.sql); i++ {
		if matched, _ := regexp.MatchString(`[a-zA-Z0-9_*]`, string(p.sql[i])); !matched {
			return p.sql[p.i:i], len(p.sql[p.i:i])
		}
	}
	return p.sql[p.i:], len(p.sql[p.i:])
}

func (p *parser) validate() error {
	if len(p.query.Conditions) == 0 && p.step == stepWhereField {
		return fmt.Errorf("at WHERE: empty WHERE clause")
	}
	if p.query.Type == query.UnknownType {
		return fmt.Errorf("query type cannot be empty")
	}
	if p.query.TableName == "" {
		return fmt.Errorf("table name cannot be empty")
	}
	if len(p.query.Conditions) == 0 && (p.query.Type == query.Update || p.query.Type == query.Delete) {
		return fmt.Errorf("at WHERE: WHERE clause is mandatory for UPDATE & DELETE")
	}
	for _, c := range p.query.Conditions {
		if c.Operator == query.UnknownOperator {
			return fmt.Errorf("at WHERE: condition without operator")
		}
		if c.Operand1 == "" && c.Operand1IsField {
			return fmt.Errorf("at WHERE: condition with empty left side operand")
		}
		if c.Operand2 == "" && c.Operand2IsField {
			return fmt.Errorf("at WHERE: condition with empty right side operand")
		}
	}
	if p.query.Type == query.Insert && len(p.query.Inserts) == 0 {
		return fmt.Errorf("at INSERT INTO: need at least one row to insert")
	}
	if p.query.Type == query.Insert {
		for _, i := range p.query.Inserts {
			if len(i) != len(p.query.Fields) {
				return fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}
		}
	}
	return nil
}

func (p *parser) logError() {
	if p.err == nil {
		return
	}
	fmt.Println(p.sql)
	fmt.Println(strings.Repeat(" ", p.i) + "^")
	fmt.Println(p.err)
}

func isIdentifier(s string) bool {
	// TODO 可以优化  使用map进行优化
	// 在关键词中的， 是不合法的
	for _, rw := range reservedWords {
		if strings.ToUpper(s) == rw {
			return false
		}
	}
	// 只能是数字和字母 是合法的
	matched, _ := regexp.MatchString("[a-zA-Z_][a-zA-Z_0-9]*", s)
	return matched
}

func isIdentifierOrAsterisk(s string) bool {
	return isIdentifier(s) || s == "*"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
