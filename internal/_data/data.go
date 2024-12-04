// package that includes mock values to be used in testing
package _data

var (
	MockEntityValid1  = []byte(`{"id": 1, "name": "TestEntity1"}`)
	MockEntityValid2  = []byte(`{"id": 1, "name": "TestEntity2"}`)
	MockEntityInvalid = []byte(`{"id": 1, "name": "TestEnti`)
)
