//memo包提供了一个对类型Func并发不安全的函数记忆功能
//函数记忆：即缓存函数的结果，达到多次调用但只需计算一次的效果
package memo1

//Func是用于记忆的函数类型
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//Memo缓存了调用Func的结果
type Memo struct {
	f Func
	cache map[string]result
}

//注意：非并发安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
