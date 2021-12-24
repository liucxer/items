package models

import (
	fmt "fmt"
	time "time"

	git_querycap_com_tools_datatypes "git.querycap.com/tools/datatypes"
	github_com_go_courier_sqlx_v2 "github.com/go-courier/sqlx/v2"
	github_com_go_courier_sqlx_v2_builder "github.com/go-courier/sqlx/v2/builder"
	github_com_go_courier_sqlx_v2_datatypes "github.com/go-courier/sqlx/v2/datatypes"
)

func (Res) PrimaryKey() []string {
	return []string{
		"ID",
	}
}

func (Res) UniqueIndexUIMd5() string {
	return "ui_md5"
}

func (Res) UniqueIndexUIResID() string {
	return "ui_res_id"
}

func (Res) UniqueIndexes() github_com_go_courier_sqlx_v2_builder.Indexes {
	return github_com_go_courier_sqlx_v2_builder.Indexes{
		"ui_md5": []string{
			"Md5",
		},
		"ui_res_id": []string{
			"ResID",
		},
	}
}

func (Res) Comments() map[string]string {
	return map[string]string{
		"CreatedAt": "创建时间",
		"Filename":  "文件名用于回显资源",
		"Info":      "资源描述信息 用于扩展资源信息 如版本号",
		"Md5":       "资源MD5",
		"ResID":     "资源ID",
		"Type":      "资源类型",
		"UpdatedAt": "更新时间",
	}
}

var ResTable *github_com_go_courier_sqlx_v2_builder.Table

func init() {
	ResTable = DB.Register(&Res{})
}

type ResIterator struct {
}

func (ResIterator) New() interface{} {
	return &Res{}
}

func (ResIterator) Resolve(v interface{}) *Res {
	return v.(*Res)
}

func (Res) TableName() string {
	return "t_res"
}

func (Res) TableDescription() []string {
	return []string{
		"Res 静态资源",
	}
}

func (Res) ColDescriptions() map[string][]string {
	return map[string][]string{
		"CreatedAt": []string{
			"创建时间",
		},
		"Filename": []string{
			"文件名用于回显资源",
		},
		"Info": []string{
			"资源描述信息 用于扩展资源信息 如版本号",
		},
		"Md5": []string{
			"资源MD5",
		},
		"ResID": []string{
			"资源ID",
		},
		"Type": []string{
			"资源类型",
		},
		"UpdatedAt": []string{
			"更新时间",
		},
	}
}

func (Res) FieldKeyID() string {
	return "ID"
}

func (m *Res) FieldID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyID())
}

func (Res) FieldKeyResID() string {
	return "ResID"
}

func (m *Res) FieldResID() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyResID())
}

func (Res) FieldKeyType() string {
	return "Type"
}

func (m *Res) FieldType() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyType())
}

func (Res) FieldKeyInfo() string {
	return "Info"
}

func (m *Res) FieldInfo() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyInfo())
}

func (Res) FieldKeyFilename() string {
	return "Filename"
}

func (m *Res) FieldFilename() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyFilename())
}

func (Res) FieldKeyMd5() string {
	return "Md5"
}

func (m *Res) FieldMd5() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyMd5())
}

func (Res) FieldKeyCreatedAt() string {
	return "CreatedAt"
}

func (m *Res) FieldCreatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyCreatedAt())
}

func (Res) FieldKeyUpdatedAt() string {
	return "UpdatedAt"
}

func (m *Res) FieldUpdatedAt() *github_com_go_courier_sqlx_v2_builder.Column {
	return ResTable.F(m.FieldKeyUpdatedAt())
}

func (Res) ColRelations() map[string][]string {
	return map[string][]string{}
}

func (m *Res) IndexFieldNames() []string {
	return []string{
		"ID",
		"Md5",
		"ResID",
	}
}

func (m *Res) ConditionByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) github_com_go_courier_sqlx_v2_builder.SqlCondition {
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

func (m *Res) Create(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	if m.CreatedAt.IsZero() {
		m.CreatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	if m.UpdatedAt.IsZero() {
		m.UpdatedAt = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	_, err := db.ExecExpr(github_com_go_courier_sqlx_v2.InsertToDB(db, m, nil))
	return err

}

func (m *Res) CreateOnDuplicateWithUpdateFields(db github_com_go_courier_sqlx_v2.DBExecutor, updateFields []string) error {

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

func (m *Res) DeleteByStruct(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(m.ConditionByStruct(db)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.DeleteByStruct"),
			),
	)

	return err
}

func (m *Res) FetchByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByID"),
			),
		m,
	)

	return err
}

func (m *Res) UpdateByIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

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
				github_com_go_courier_sqlx_v2_builder.Comment("Res.UpdateByIDWithMap"),
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

func (m *Res) UpdateByIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByIDWithMap(db, fieldValues)

}

func (m *Res) FetchByIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Res) DeleteByID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ID").Eq(m.ID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.DeleteByID"),
			))

	return err
}

func (m *Res) FetchByResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByResID"),
			),
		m,
	)

	return err
}

func (m *Res) UpdateByResIDWithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.UpdateByResIDWithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByResID(db)
	}

	return nil

}

func (m *Res) UpdateByResIDWithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByResIDWithMap(db, fieldValues)

}

func (m *Res) FetchByResIDForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByResIDForUpdate"),
			),
		m,
	)

	return err
}

func (m *Res) DeleteByResID(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("ResID").Eq(m.ResID),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.DeleteByResID"),
			))

	return err
}

func (m *Res) FetchByMd5(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Md5").Eq(m.Md5),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByMd5"),
			),
		m,
	)

	return err
}

func (m *Res) UpdateByMd5WithMap(db github_com_go_courier_sqlx_v2.DBExecutor, fieldValues github_com_go_courier_sqlx_v2_builder.FieldValues) error {

	if _, ok := fieldValues["UpdatedAt"]; !ok {
		fieldValues["UpdatedAt"] = github_com_go_courier_sqlx_v2_datatypes.Timestamp(time.Now())
	}

	table := db.T(m)

	result, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Update(db.T(m)).
			Where(
				github_com_go_courier_sqlx_v2_builder.And(
					table.F("Md5").Eq(m.Md5),
				),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.UpdateByMd5WithMap"),
			).
			Set(table.AssignmentsByFieldValues(fieldValues)...),
	)

	if err != nil {
		return err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return m.FetchByMd5(db)
	}

	return nil

}

func (m *Res) UpdateByMd5WithStruct(db github_com_go_courier_sqlx_v2.DBExecutor, zeroFields ...string) error {

	fieldValues := github_com_go_courier_sqlx_v2_builder.FieldValuesFromStructByNonZero(m, zeroFields...)
	return m.UpdateByMd5WithMap(db, fieldValues)

}

func (m *Res) FetchByMd5ForUpdate(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	err := db.QueryExprAndScan(
		github_com_go_courier_sqlx_v2_builder.Select(nil).
			From(
				db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Md5").Eq(m.Md5),
				)),
				github_com_go_courier_sqlx_v2_builder.ForUpdate(),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.FetchByMd5ForUpdate"),
			),
		m,
	)

	return err
}

func (m *Res) DeleteByMd5(db github_com_go_courier_sqlx_v2.DBExecutor) error {

	table := db.T(m)

	_, err := db.ExecExpr(
		github_com_go_courier_sqlx_v2_builder.Delete().
			From(db.T(m),
				github_com_go_courier_sqlx_v2_builder.Where(github_com_go_courier_sqlx_v2_builder.And(
					table.F("Md5").Eq(m.Md5),
				)),
				github_com_go_courier_sqlx_v2_builder.Comment("Res.DeleteByMd5"),
			))

	return err
}

func (m *Res) List(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) ([]Res, error) {

	list := make([]Res, 0)

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Res.List"),
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

func (m *Res) Count(db github_com_go_courier_sqlx_v2.DBExecutor, condition github_com_go_courier_sqlx_v2_builder.SqlCondition, additions ...github_com_go_courier_sqlx_v2_builder.Addition) (int, error) {

	count := -1

	table := db.T(m)
	_ = table

	finalAdditions := []github_com_go_courier_sqlx_v2_builder.Addition{
		github_com_go_courier_sqlx_v2_builder.Where(condition),
		github_com_go_courier_sqlx_v2_builder.Comment("Res.Count"),
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

func (m *Res) BatchFetchByIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []uint64) ([]Res, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ID").In(values)

	return m.List(db, condition)

}

func (m *Res) BatchFetchByMd5List(db github_com_go_courier_sqlx_v2.DBExecutor, values []string) ([]Res, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("Md5").In(values)

	return m.List(db, condition)

}

func (m *Res) BatchFetchByResIDList(db github_com_go_courier_sqlx_v2.DBExecutor, values []git_querycap_com_tools_datatypes.SFID) ([]Res, error) {

	if len(values) == 0 {
		return nil, nil
	}

	table := db.T(m)

	condition := table.F("ResID").In(values)

	return m.List(db, condition)

}
