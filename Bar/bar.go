package bar

import foo "goAvancado/Foo"

func TakeFoo(i foo.Interface) {
	i.Interface()
}
