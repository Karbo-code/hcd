package mempool

import (
	"github.com/HcashOrg/hcd/hcutil"
)

type addrPool struct {
	addrPool map[string]interface{}
}


func (mp *TxPool) AddToAddrPool(addr string) error {
	mp.mtx.Lock()
	defer mp.mtx.Unlock()
	return mp.maybeAddtoAddrPool(addr)
}

func (mp *TxPool) maybeAddtoAddrPool(addr string) error {
	_,err:=hcutil.DecodeAddress(addr)
	if err!=nil{
		return err
	}
	mp.addrPool.addrPool[addr] = nil

	return nil
}



func (mp *TxPool) RemoveAddr(addr string) {
	mp.mtx.Lock()
	defer mp.mtx.Unlock()

	mp.removeAddr(addr)
}

func (mp *TxPool) removeAddr(addr string) {
	delete(mp.addrPool.addrPool,addr)
}

func (mp *TxPool)GetAddrList()[]string{
	mp.mtx.RLock()
	defer mp.mtx.RUnlock()
	return mp.getAddrList()
}

func (mp *TxPool)getAddrList()[]string {

	addrSlice:=make([]string,0,len(mp.addrPool.addrPool))

	for addr,_:=range mp.addrPool.addrPool {
		addrSlice=append(addrSlice,addr)
	}
	return addrSlice
}