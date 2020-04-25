package main
import "github.com/nareshkumarthota/flogocomponents/activity/methodinvoker"
func kafkaMethod(inputs interface{}) (map[string]interface{}, error) {
	return nil, nil
}
func init() {
	methodinvoker.RegisterMethods("kafkaMethod", kafkaMethod)
}
