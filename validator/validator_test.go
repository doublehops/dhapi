package validator

//type Car struct {
//	Make  string
//	Model string
//	Year  int
//}
//
//func TestValidate(t *testing.T) {
//	tests := []struct {
//		Name           string
//		Rules          []Rule
//		Payload        Car
//		expectedErrors []ErrorMessage
//	}{
//		{
//			Name: "AllPass",
//			Rules: []Rule{
//				{
//					VariableName: "Name",
//					Required:     true,
//					Function:     Min(1, 2)(1, 2),
//				},
//			},
//		},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.Name, func(t *testing.T) {
//
//			v := []Rule{}
//			val := Min()(1, 2)
//			a := val(1, 2)
//
//			//for _, rule := range tt.Rules {
//			//	rule.Value = tt.Payload.
//			//	{
//			//		rule.VariableName
//			//	}
//			//}
//			//errors := Validate("tt.Payload", tt.Rules)
//			//if !reflect.DeepEqual(tt.expectedErrors, errors) {
//			//	t.Errorf("Unexpected validation response. Expected: %v; Actual: %v", tt.expectedErrors, errors)
//			//}
//		})
//	}
//}
//
//func validations(car Car) []Rule {
//	return []Rule{
//		{"make", car.Make, true, NotEmptyAAA("asd", true)},
//	}
//}
//
//func (c *Car) Validate() []ErrorMessage {
//	var errors []ErrorMessage
//
//	return errors
//}
