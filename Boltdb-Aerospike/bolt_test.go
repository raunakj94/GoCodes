import "testing"
func BenchmarkUpdate(b *testing.B) {
    fn1 := Update()
    for i := 0; i < b.N; i++ {
        _ = fn1()
    }
func BenchmarkView(b *testing.B) {
    fn := View()
    for i := 0; i < b.N; i++ {
        _ = fn()
    }
~       
