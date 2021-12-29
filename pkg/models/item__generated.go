package models

import (
	fmt "fmt"
	time "time"

	git_querycap_com_tools_datatypes "git.querycap.com/tools/datatypes"
	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
)

func (Item) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Item) UniqueIndexUICode() string {
	return "ui_code"
}

func (Item) UniqueIndexUICodeParentCode() string {
	return "ui_code_parent_code"
}

func (Item) UniqueIndexUIItemID() string {
	return "ui_item_id"
}

func (Item) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"ui_code": []string{
			"Code",
		},
		"ui_code_parent_code": []string{
			"Code",
			"ParentCode",
		},
		"ui_item_id": []string{
			"ItemID",
		},
	}
}

func (Item) Comments() map[string]string {
	return map[string]string{
		"AlphabetEN":  "条目英文",
		"AlphabetZH":  "条目中文拼音",
		"AttachResID": "条目附件资源ID",
		"Code":        "条目代码",
		"CreatedAt":   "创建时间",
		"HasSub":      "是否有下一级条目",
		"ImageResID":  "条目icon资源ID",
		"ItemID":      "条目编号(UUID)",
		"Link":        "条目跳转链接",
		"Name":        "条目名称",
		"ParentCode":  "条目上级代码",
		"RichText":    "条目富文本",
		"UpdatedAt":   "更新时间",
	}
}

var ItemTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	ItemTable = DB.Register(&Item{})
}

type ItemIterator struct {
}

func (ItemIterator) New() interface{} {
	return &Item{}
}

func (ItemIterator) Resolve(v interface{}) *Item {
	return v.(*Item)
}

func (Item) TableName() string {
	return "t_item"
}

func (Item) TableDescription() []string {
	return []string{
		"Item 条目信息",
	}
}

func (Item) ColDescriptions() map[string][]string {
	return map[string][]string{
		"AlphabetEN": []string{
			"条目英文",
		},
		"AlphabetZH": []string{
			"条目中文拼音",
		},
		"AttachResID": []string{
			"条目附件资源ID",
		},
		"Code": []string{
			"条目代码",
		},
		"CreatedAt": []string{
			"创建时间",
		},
		"HasSub": []string{
			"是否有下一级条目",
		},
		"ImageResID": []string{
			"条目icon资源ID",
		},
		"ItemID": []string{
			"条目编号(UUID)",
		},
		"Link": []string{
			"条目跳转链接",
		},
		"Name": []string{
			"条目名称",
		},
		"ParentCode": []string{
			"条目上级代码",
		},
		"RichText": []string{
			"条目富文本",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
	}
}

func (Item) FieldKeyID() string {
	return "ID"
}

func (m *Item) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyID())
}

func (Item) FieldKeyItemID() string {
	return "ItemID"
}

func (m *Item) FieldItemID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyItemID())
}

func (Item) FieldKeyCode() string {
	return "Code"
}

func (m *Item) FieldCode() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyCode())
}

func (Item) FieldKeyParentCode() string {
	return "ParentCode"
}

func (m *Item) FieldParentCode() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyParentCode())
}

func (Item) FieldKeyName() string {
	return "Name"
}

func (m *Item) FieldName() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyName())
}

func (Item) FieldKeyAlphabetZH() string {
	return "AlphabetZH"
}

func (m *Item) FieldAlphabetZH() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyAlphabetZH())
}

func (Item) FieldKeyAlphabetEN() string {
	return "AlphabetEN"
}

func (m *Item) FieldAlphabetEN() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyAlphabetEN())
}

func (Item) FieldKeyImageResID() string {
	return "ImageResID"
}

func (m *Item) FieldImageResID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyImageResID())
}

func (Item) FieldKeyRichText() string {
	return "RichText"
}

func (m *Item) FieldRichText() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyRichText())
}

func (Item) FieldKeyLink() string {
	return "Link"
}

func (m *Item) FieldLink() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyLink())
}

func (Item) FieldKeyAttachResID() string {
	return "AttachResID"
}

func (m *Item) FieldAttachResID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyAttachResID())
}

func (Item) FieldKeyHasSub() string {
	return "HasSub"
}

func (m *Item) FieldHasSub() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyHasSub())
}

func (Item) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Item) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyCreatedAt())
}

func (Item) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Item) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ItemTable.F(m.FieldKeyUpdatedAt())
}

func (Item) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Item) IndexFieldNames() []string {
	return []string{
		"Code",
		"ID",
		"ItemID",
		"ParentCode",
	}
}

func (m *Item) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
	table := db.T(m)
	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m)

	conditions := make([]github_com_go_courier_sqlx_v2_builder.SqlCondition, 0)

	for _, fieldName := range m.IndexFieldNames() {
		if v, exists := fieldValues[fieldName]; exists {
			conditions = append(conditions, table.F(fieldName).Eq(v))
			delete(fieldValues, fieldName)
		}
	}

	if len(conditions) == 0 {
		panic(fmt.Errorf("at least one of field for indexes has value"))
	}

	for fieldName, v := range fieldValues {
		conditions = append(conditions, table.F(fieldName).Eq(v))
	}

	condition := github_com_go_courier_sqlx_v2_builder.And(conditions...)

	return condition
}

func (m *Item) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Item) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

	if len(updateFields) == 0 {
		panic(fmt.Errorf("must have update fields"))
	}

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, updateFields...)

	delete(fieldValues, "ID")

	table := db.T(m)

	cols, vals := table.ColumnsAndValuesByFieldValues(fieldValues)

	fields := make(map[string]bool, len(updateFields))
	for _, field := range updateFields {
		fields[field] = true
	}

	for _, fieldNames := range m.UniqueIndexes() {
		for _, field := range fieldNames {
			delete(fields, field)
		}
	}

	if len(fields) == 0 {
		panic(fmt.Errorf("no fields for updates"))
	}

	for field := range fieldValues {
		if !fields[field] {
			delete(fieldValues, field)
		}
	}

	additions := github_com_go_courier_sqlx_v2_builder.Additions{}

	switch db.Dialect().DriverName() {
	case "mysql":
		additions = append(additions, github_com_go_courier_sqlx_v2_builder.OnDuplicateKeyUpdate(table.AssignmentsByFieldValues(fieldValues)...))
	case "postgres":
		indexes := m.UniqueIndexes()
		fields := make([]string, 0)
		for _, fs := range indexes {
			fields = append(fields, fs...)
		}
		indexFields, _ := db.T(m).Fields(fields...)

		additions = append(additions,
			github_com_go_courier_sqlx_v2_builder.OnConflict(indexFields).
				DoUpdateSet(table.AssignmentsByFieldValues(fieldValues)...))
	}

	additions = append(additions, github_com_go_courier_sqlx_v2_builder.Comment("User.CreateOnDuplicateWithUpdateFields"))

	expr := github_com_go_courier_sqlx_v2_builder.Insert().Into(table, additions...).Values(cols, vals...)

	_, err := db.ExecExpr(expr)
	return err

}

func (m *Item) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.DeleteByStruct"),
			),
	)

	return err
}

func (m *Item) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Item) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.UpdateByIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByID(db)
	}

	return nil

}

func (m *Item) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Item) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Item) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.DeleteByID"),
			))

	return err
}

func (m *Item) FetchByItemID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ItemID").Eq(m.ItemID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByItemID"),
			),
		m,
	)

	return err
}

func (m *Item) UpdateByItemIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ItemID").Eq(m.ItemID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.UpdateByItemIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByItemID(db)
	}

	return nil

}

func (m *Item) UpdateByItemIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByItemIDWithMap(db, fieldValues)

}

func (m *Item) FetchByItemIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ItemID").Eq(m.ItemID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByItemIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Item) DeleteByItemID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ItemID").Eq(m.ItemID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.DeleteByItemID"),
			))

	return err
}

func (m *Item) FetchByCode(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByCode"),
			),
		m,
	)

	return err
}

func (m *Item) UpdateByCodeWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.UpdateByCodeWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByCode(db)
	}

	return nil

}

func (m *Item) UpdateByCodeWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByCodeWithMap(db, fieldValues)

}

func (m *Item) FetchByCodeForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByCodeForUpdate"),
			),
		m,
	)

	return err
}

func (m *Item) DeleteByCode(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.DeleteByCode"),
			))

	return err
}

func (m *Item) FetchByCodeAndParentCode(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
					table.F("ParentCode").Eq(m.ParentCode),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByCodeAndParentCode"),
			),
		m,
	)

	return err
}

func (m *Item) UpdateByCodeAndParentCodeWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
					table.F("ParentCode").Eq(m.ParentCode),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.UpdateByCodeAndParentCodeWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByCodeAndParentCode(db)
	}

	return nil

}

func (m *Item) UpdateByCodeAndParentCodeWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByCodeAndParentCodeWithMap(db, fieldValues)

}

func (m *Item) FetchByCodeAndParentCodeForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
					table.F("ParentCode").Eq(m.ParentCode),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.FetchByCodeAndParentCodeForUpdate"),
			),
		m,
	)

	return err
}

func (m *Item) DeleteByCodeAndParentCode(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Code").Eq(m.Code),
					table.F("ParentCode").Eq(m.ParentCode),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Item.DeleteByCodeAndParentCode"),
			))

	return err
}

func (m *Item) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Item, error) {

	list := make([]Item, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Item.List"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(db.T(m), finalAdditions...),
		&list,
	)

	return list, err

}

func (m *Item) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Item.Count"),
	}

	if len(additions) > 0 {
		finalAdditions = append(finalAdditions, additions...)
	}

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(
			github_com_go_courier_sqlx_v2_builder.Count(),
		).
			From(db.T(m), finalAdditions...),
		&count,
	)

	return count, err

}

func (m *Item) BatchFetchByCodeList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Item, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Code").In(values)

	return m.List(db, condition)

}

func (m *Item) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Item, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Item) BatchFetchByItemIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]Item, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ItemID").In(values)

	return m.List(db, condition)

}

func (m *Item) BatchFetchByParentCodeList(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Item, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ParentCode").In(values)

	return m.List(db, condition)

}
