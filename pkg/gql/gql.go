package gql

import (
	"errors"
	"fmt"
	"reflect"
)

func NewObject(Name string, Fields string) string {
	return fmt.Sprintf(`
var %s =  graphql.NewObject(graphql.ObjectConfig{
	Name: "%s",
	Fields: graphql.Fields{
		%s
	},
})
`, Name, Name, Fields)
}

func Field(Name string, Type string, c string) (string, error) {
	if Name == "Password" || Name == "Salt" {
		return "", nil
	}
	//log.Println(Name, Type)
	switch Type {
	case "uuid.UUID":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.String}`+c, Name), nil

	case "*models.Board":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: Board}`+c, Name), nil

	case "*models.Post":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: Post}`+c, Name), nil

	case "*models.User":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: User}`+c, Name), nil

	case "int":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.Int}`+c, Name), nil
	//todo depricated int64 != GraphQLInt
	case "sql.NullInt64":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.Int}`+c, Name), nil
	//todo depricated int64 != GraphQLInt
	case "int64":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.Int}`+c, Name), nil
	case "pq.NullTime":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.DateTime}`+c, Name), nil
	case "time.Time":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.DateTime}`+c, Name), nil
	case "bool":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.%s}`+c, Name, "Boolean"), nil
	case "sql.NullString":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.String}`+c, Name), nil
	case "string":
		return fmt.Sprintf(`"%s": &graphql.Field{ Type: graphql.String}`+c, Name), nil
	}
	return "", errors.New("unknown type: " + Type)
}

func TypeToGqlType(Name string, i interface{}) (res string, err error) {
	t := reflect.TypeOf(i)
	//log.Print(t)
	comma := ",\n\t\t"
	fields := ""
	for i := 0; i < t.NumField(); i++ {
		//log.Print(t.Field(i))
		field, err := tpl(t.Field(i), comma)
		if err != nil {
			return "", err
		}
		fields = fields + field
	}
	res = fmt.Sprint(NewObject(Name, fields))
	return
}

func tpl(f reflect.StructField, comma string) (res string, err error) {
	return Field(f.Name, f.Type.String(), comma)
}
