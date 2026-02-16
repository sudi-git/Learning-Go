package sprint 
func IntVsFloat(i int, f float32) string {
in := float32 (i)
if in > f {
    return "Integer"
} else if f > in {
    return "Float" 
} else {
    return "Same"
}
}