package exp

type (
	SelectClauses interface {
		HasSources() bool
		IsDefaultSelect() bool
		clone() *selectClauses
		Clone() SelectClauses

		Select() ColumnListExpression
		SelectAppend(cl ColumnListExpression) SelectClauses
		SetSelect(cl ColumnListExpression) SelectClauses

		Distinct() ColumnListExpression
		SetDistinct(cle ColumnListExpression) SelectClauses

		From() ColumnListExpression
		SetFrom(cl ColumnListExpression) SelectClauses

		HasAlias() bool
		Alias() IdentifierExpression
		SetAlias(ie IdentifierExpression) SelectClauses

		Joins() JoinExpressions
		JoinsAppend(jc JoinExpression) SelectClauses

		AsOfSystemTime() interface{}
		HasAsOfSystemTime() bool
		ClearAsOfSystemTime() SelectClauses
		SetAsOfSystemTime(limit interface{}) SelectClauses

		Where() ExpressionList
		ClearWhere() SelectClauses
		WhereAppend(expressions ...Expression) SelectClauses

		Having() ExpressionList
		ClearHaving() SelectClauses
		HavingAppend(expressions ...Expression) SelectClauses

		Order() ColumnListExpression
		HasOrder() bool
		ClearOrder() SelectClauses
		SetOrder(oes ...OrderedExpression) SelectClauses
		OrderAppend(...OrderedExpression) SelectClauses
		OrderPrepend(...OrderedExpression) SelectClauses

		GroupBy() ColumnListExpression
		SetGroupBy(cl ColumnListExpression) SelectClauses

		Limit() interface{}
		HasLimit() bool
		ClearLimit() SelectClauses
		SetLimit(limit interface{}) SelectClauses

		Offset() uint
		ClearOffset() SelectClauses
		SetOffset(offset uint) SelectClauses

		Compounds() []CompoundExpression
		CompoundsAppend(ce CompoundExpression) SelectClauses

		Lock() Lock
		SetLock(l Lock) SelectClauses

		CommonTables() []CommonTableExpression
		CommonTablesAppend(cte CommonTableExpression) SelectClauses

		Windows() []WindowExpression
		SetWindows(ws []WindowExpression) SelectClauses
		WindowsAppend(ws ...WindowExpression) SelectClauses
		ClearWindows() SelectClauses
	}
	selectClauses struct {
		commonTables   []CommonTableExpression
		selectColumns  ColumnListExpression
		distinct       ColumnListExpression
		from           ColumnListExpression
		joins          JoinExpressions
		asOfSystemTime interface{}
		where          ExpressionList
		alias          IdentifierExpression
		groupBy        ColumnListExpression
		having         ExpressionList
		order          ColumnListExpression
		limit          interface{}
		offset         uint
		compounds      []CompoundExpression
		lock           Lock
		windows        []WindowExpression
	}
)

func NewSelectClauses() SelectClauses {
	return &selectClauses{
		selectColumns: NewColumnListExpression(Star()),
	}
}

func (c *selectClauses) HasSources() bool {
	return c.from != nil && len(c.from.Columns()) > 0
}

func (c *selectClauses) IsDefaultSelect() bool {
	ret := false
	if c.selectColumns != nil {
		selects := c.selectColumns.Columns()
		if len(selects) == 1 {
			if l, ok := selects[0].(LiteralExpression); ok && l.Literal() == "*" {
				ret = true
			}
		}
	}
	return ret
}

func (c *selectClauses) clone() *selectClauses {
	var clone selectClauses
	if len(c.commonTables) > 0 {
		clone.commonTables = make([]CommonTableExpression, len(c.commonTables))
		copy(clone.commonTables, c.commonTables)
	}

	if c.selectColumns != nil {
		clone.selectColumns = c.selectColumns.Clone().(ColumnListExpression)
	}

	if c.distinct != nil {
		clone.distinct = c.distinct.Clone().(ColumnListExpression)
	}

	if c.distinct != nil {
		clone.distinct = c.distinct.Clone().(ColumnListExpression)
	}

	if c.from != nil {
		clone.from = c.from.Clone().(ColumnListExpression)
	}

	if len(c.joins) > 0 {
		clone.joins = make(JoinExpressions, len(c.joins))
		copy(clone.joins, c.joins)
	}

	clone.asOfSystemTime = c.asOfSystemTime

	if c.where != nil {
		clone.where = c.where.Clone().(ExpressionList)
	}

	if c.alias != nil {
		clone.alias = c.alias.Clone().(IdentifierExpression)
	}

	if c.groupBy != nil {
		clone.groupBy = c.groupBy.Clone().(ColumnListExpression)
	}

	if c.having != nil {
		clone.having = c.having.Clone().(ExpressionList)
	}

	if c.order != nil {
		clone.order = c.order.Clone().(ColumnListExpression)
	}

	clone.limit = c.limit
	clone.offset = c.offset

	if len(c.compounds) > 0 {
		clone.compounds = make([]CompoundExpression, len(c.compounds))
		copy(clone.compounds, c.compounds)
	}

	if len(c.windows) > 0 {
		clone.windows = make([]WindowExpression, len(c.windows))
		copy(clone.windows, c.windows)
	}

	clone.lock = c.lock

	return &clone
}

func (c *selectClauses) Clone() SelectClauses {
	return c.clone()
}

func (c *selectClauses) CommonTables() []CommonTableExpression {
	return c.commonTables
}

func (c *selectClauses) CommonTablesAppend(cte CommonTableExpression) SelectClauses {
	ret := c.clone()
	ret.commonTables = append(ret.commonTables, cte)
	return ret
}

func (c *selectClauses) Select() ColumnListExpression {
	return c.selectColumns
}

func (c *selectClauses) SelectAppend(cl ColumnListExpression) SelectClauses {
	ret := c.clone()
	ret.selectColumns = ret.selectColumns.Append(cl.Columns()...)
	return ret
}

func (c *selectClauses) SetSelect(cl ColumnListExpression) SelectClauses {
	ret := c.clone()
	ret.selectColumns = cl
	return ret
}

func (c *selectClauses) Distinct() ColumnListExpression {
	return c.distinct
}

func (c *selectClauses) SetDistinct(cle ColumnListExpression) SelectClauses {
	ret := c.clone()
	ret.distinct = cle
	return ret
}

func (c *selectClauses) From() ColumnListExpression {
	return c.from
}

func (c *selectClauses) SetFrom(cl ColumnListExpression) SelectClauses {
	ret := c.clone()
	ret.from = cl
	return ret
}

func (c *selectClauses) HasAlias() bool {
	return c.alias != nil
}

func (c *selectClauses) Alias() IdentifierExpression {
	return c.alias
}

func (c *selectClauses) SetAlias(ie IdentifierExpression) SelectClauses {
	ret := c.clone()
	ret.alias = ie
	return ret
}

func (c *selectClauses) Joins() JoinExpressions {
	return c.joins
}

func (c *selectClauses) JoinsAppend(jc JoinExpression) SelectClauses {
	ret := c.clone()
	ret.joins = append(ret.joins, jc)
	return ret
}

func (c *selectClauses) AsOfSystemTime() interface{} {
	return c.asOfSystemTime
}

func (c *selectClauses) HasAsOfSystemTime() bool {
	return c.asOfSystemTime != nil
}

func (c *selectClauses) ClearAsOfSystemTime() SelectClauses {
	ret := c.clone()
	ret.asOfSystemTime = nil
	return ret
}

func (c *selectClauses) SetAsOfSystemTime(asOfSystemTime interface{}) SelectClauses {
	ret := c.clone()
	ret.asOfSystemTime = asOfSystemTime
	return ret
}

func (c *selectClauses) Where() ExpressionList {
	return c.where
}

func (c *selectClauses) ClearWhere() SelectClauses {
	ret := c.clone()
	ret.where = nil
	return ret
}

func (c *selectClauses) WhereAppend(expressions ...Expression) SelectClauses {
	if len(expressions) == 0 {
		return c
	}
	ret := c.clone()
	if ret.where == nil {
		ret.where = NewExpressionList(AndType, expressions...)
	} else {
		ret.where = ret.where.Append(expressions...)
	}
	return ret
}

func (c *selectClauses) Having() ExpressionList {
	return c.having
}

func (c *selectClauses) ClearHaving() SelectClauses {
	ret := c.clone()
	ret.having = nil
	return ret
}

func (c *selectClauses) HavingAppend(expressions ...Expression) SelectClauses {
	if len(expressions) == 0 {
		return c
	}
	ret := c.clone()
	if ret.having == nil {
		ret.having = NewExpressionList(AndType, expressions...)
	} else {
		ret.having = ret.having.Append(expressions...)
	}
	return ret
}

func (c *selectClauses) Lock() Lock {
	return c.lock
}

func (c *selectClauses) SetLock(l Lock) SelectClauses {
	ret := c.clone()
	ret.lock = l
	return ret
}

func (c *selectClauses) Order() ColumnListExpression {
	return c.order
}

func (c *selectClauses) HasOrder() bool {
	return c.order != nil
}

func (c *selectClauses) ClearOrder() SelectClauses {
	ret := c.clone()
	ret.order = nil
	return ret
}

func (c *selectClauses) SetOrder(oes ...OrderedExpression) SelectClauses {
	ret := c.clone()
	ret.order = NewOrderedColumnList(oes...)
	return ret
}

func (c *selectClauses) OrderAppend(oes ...OrderedExpression) SelectClauses {
	if c.order == nil {
		return c.SetOrder(oes...)
	}
	ret := c.clone()
	ret.order = ret.order.Append(NewOrderedColumnList(oes...).Columns()...)
	return ret
}

func (c *selectClauses) OrderPrepend(oes ...OrderedExpression) SelectClauses {
	if c.order == nil {
		return c.SetOrder(oes...)
	}
	ret := c.clone()
	ret.order = NewOrderedColumnList(oes...).Append(ret.order.Columns()...)
	return ret
}

func (c *selectClauses) GroupBy() ColumnListExpression {
	return c.groupBy
}

func (c *selectClauses) SetGroupBy(cl ColumnListExpression) SelectClauses {
	ret := c.clone()
	ret.groupBy = cl
	return ret
}

func (c *selectClauses) Limit() interface{} {
	return c.limit
}

func (c *selectClauses) HasLimit() bool {
	return c.limit != nil
}

func (c *selectClauses) ClearLimit() SelectClauses {
	ret := c.clone()
	ret.limit = nil
	return ret
}

func (c *selectClauses) SetLimit(limit interface{}) SelectClauses {
	ret := c.clone()
	ret.limit = limit
	return ret
}

func (c *selectClauses) Offset() uint {
	return c.offset
}

func (c *selectClauses) ClearOffset() SelectClauses {
	ret := c.clone()
	ret.offset = 0
	return ret
}

func (c *selectClauses) SetOffset(offset uint) SelectClauses {
	ret := c.clone()
	ret.offset = offset
	return ret
}

func (c *selectClauses) Compounds() []CompoundExpression {
	return c.compounds
}

func (c *selectClauses) CompoundsAppend(ce CompoundExpression) SelectClauses {
	ret := c.clone()
	ret.compounds = append(ret.compounds, ce)
	return ret
}

func (c *selectClauses) Windows() []WindowExpression {
	return c.windows
}

func (c *selectClauses) SetWindows(ws []WindowExpression) SelectClauses {
	ret := c.clone()
	ret.windows = ws
	return ret
}

func (c *selectClauses) WindowsAppend(ws ...WindowExpression) SelectClauses {
	ret := c.clone()
	ret.windows = append(ret.windows, ws...)
	return ret
}

func (c *selectClauses) ClearWindows() SelectClauses {
	ret := c.clone()
	ret.windows = nil
	return ret
}
