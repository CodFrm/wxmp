package model

type Answer struct {
	Index  int32
	Result []struct {
		Correct []struct {
			Content interface{}
			Option  interface{}
		}
		Type int32
	}
}
