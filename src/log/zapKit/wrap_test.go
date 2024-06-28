package zapKit

//func TestWrapLogger(t *testing.T) {
//	f, err := os.Create("_a.log")
//	if err != nil {
//		panic(err)
//	}
//
//	tmp := NewLogger(WithWriter(f))
//	l := WrapLogger(tmp, f)
//
//	l.Debug("debug")
//	l.Info("info")
//	l.Warn("warn")
//	l.Error("error")
//
//	fmt.Println("close", l.Close())
//
//	fmt.Println("=== sleep starts ===")
//	time.Sleep(time.Second * 3)
//	fmt.Println("=== sleep ends ===")
//}
