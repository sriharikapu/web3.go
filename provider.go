package web3

func (w3 *Web3) CurrentProvider() bool {
	return w3.js.Get("currentProvider") == nil
}

func (w3 *Web3) SetProvider(args ...interface{}) {
	w3.js.Call("setProvider",
		w3.js.
			Get("providers").
			Get("HttpProvider").
			New(args))

}
