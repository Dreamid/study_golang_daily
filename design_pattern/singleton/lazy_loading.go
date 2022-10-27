package singleton

import "sync"

type singletonType struct{

}

var instance * singletonType
var mu sync.Mutex

//thread safe with mutex
func GetInstanceWithMutex()*singletonType{
    mu.Lock()
	defer mu.Unlock()
	if instance == nil{
		instance = &singletonType{}
	}
	return instance
}




// not thread safe
func GetInstanceNotThreadSafe()*singletonType{
    if instance == nil{
        instance = &singletonType{}
    }
    return instance
}


