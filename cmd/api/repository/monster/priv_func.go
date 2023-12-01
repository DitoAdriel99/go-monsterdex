package monster

import (
	"fmt"
	"regexp"
	"strings"
)

type payload struct {
	args                    []any
	counter                 int
	original, limit, offset string
	where, order, groupBy   []string
}

type direction string

const (
	ASC  direction = "ASC"
	DESC direction = "DESC"
)

func (d direction) toString() string {
	return string(d)
}

type operator string

const (
	equal            operator = "="
	notEqueal        operator = "!="
	lessThan         operator = "<"
	greaterThan      operator = ">"
	lessThanEqual    operator = "<="
	greatedThanEqual operator = ">="
	between          operator = "BETWEEN"
	like             operator = "LIKE"
	in               operator = "IN"
	exists           operator = "EXISTS"
	subQuery         operator = "SUBQUERY"
	is               operator = "IS"
)

func (o operator) toString() string {
	return string(o)
}

func QueryStatement(stmt string, args ...any) *payload {
	var (
		// we can assume stmt is set properly
		counter  = 1 + strings.Count(stmt, "$")
		remove_t = strings.ReplaceAll(stmt, "\t", "")
		remove_n = strings.TrimSpace(strings.ReplaceAll(remove_t, "\n", " "))
	)

	var (
		stat = &payload{
			original: remove_n,
			counter:  counter,
			args:     args,
			where:    make([]string, 0),
			order:    make([]string, 0),
		}
	)

	return stat
}

func (p *payload) Where(by string, operator operator, isOr bool, value ...any) *payload {
	var (
		preform = ""
	)

	if len(p.where) == 0 {
		preform += "WHERE "
	} else {
		if isOr {
			preform += "OR "
		} else {
			preform += "AND "
		}
	}

	switch operator {
	case equal, notEqueal, lessThan, greaterThan, greatedThanEqual, lessThanEqual:
		if isOr {
			preform += fmt.Sprintf("%v %v ANY($%v)", by, operator.toString(), p.counter)
		} else {
			preform += fmt.Sprintf("%v %v $%v", by, operator.toString(), p.counter)

		}
		p.counter++
	case between:
		preform += fmt.Sprintf("%v %v $%v AND $%v", by, operator.toString(), p.counter, p.counter+1)
		p.counter += 2
	case like:
		preform += fmt.Sprintf("LOWER(%v::text) %v $%v", by, operator.toString(), p.counter)
		p.counter++
	case in:
		preform += fmt.Sprintf("%v %v (", by, operator.toString())
		for i := 0; i < len(value); i++ {
			preform += fmt.Sprintf("$%v", p.counter)
			if i != len(value)-1 {
				preform += ", "
			} else {
				preform += ")"
			}
			p.counter++
		}
	case subQuery, exists:
		var (
			re    = regexp.MustCompile(`[\$]+[\d]`)
			count = strings.Count(by, "$")
		)

		for i := 0; i < count; i++ {
			v := re.FindAllStringIndex(by, -1)
			by = by[:v[i][0]] + fmt.Sprintf("$%v", p.counter) + by[v[i][1]:]
			p.counter++
		}

		if operator == subQuery {
			preform += fmt.Sprintf("(%v)", by)

		} else if operator == exists {
			preform += fmt.Sprintf("%v (%v)", operator.toString(), by)

		}
	}

	p.where = append(p.where, preform)
	p.args = append(p.args, value...)

	return p
}

func (p *payload) Order(by string, direction direction) *payload {
	var (
		preform = ""
	)

	if len(p.order) == 0 {
		preform += "ORDER BY "
	} else {
		preform += ","
	}

	preform += fmt.Sprintf("%v %v", by, strings.ToUpper(direction.toString()))
	p.order = append(p.order, preform)

	return p
}

func (p *payload) Limit(limit int) *payload {
	preform := fmt.Sprintf("LIMIT $%v", p.counter)
	p.limit = preform
	p.args = append(p.args, limit)
	p.counter++

	return p
}

func (p *payload) Offset(offset int) *payload {
	preform := fmt.Sprintf("OFFSET $%v", p.counter)
	p.offset = preform
	p.args = append(p.args, offset)
	p.counter++

	return p
}

func (p *payload) Build() (string, int, []any) {
	if p == nil {
		return "", 0, make([]any, 0)
	}

	var (
		counter = p.counter
		args    = p.args
		builder = make([]string, 0)
	)

	if len(p.where) != 0 {
		builder = append(builder, strings.Join(p.where, " "))
	}

	if len(p.groupBy) != 0 {
		builder = append(builder, "GROUP BY "+strings.Join(p.groupBy, ", "))
	}
	if len(p.order) != 0 {
		builder = append(builder, strings.Join(p.order, " "))
	}

	if len(p.offset) != 0 {
		builder = append(builder, p.offset)
	}

	if len(p.limit) != 0 {
		builder = append(builder, p.limit)
	}

	build := fmt.Sprintf("%s %s;", p.original, strings.Join(builder, " "))
	return build, counter, args
}

func (p *payload) GroupBy(columns ...string) *payload {
	p.groupBy = columns
	return p
}
