type History struct {
	Url string
	Next *History
	Prev *History
}

type BrowserHistory struct {
    Current *History
}


func Constructor(homepage string) BrowserHistory {
    return BrowserHistory{Current: &History{Url: homepage}}
}


func (this *BrowserHistory) Visit(url string)  {
	if this.Current.Next != nil {
		oldUrl := this.Current.Next
		oldUrl.Prev = nil
	}
    this.Current.Next = &History{Url: url, Prev: this.Current}
	this.Current = this.Current.Next
}


func (this *BrowserHistory) Back(steps int) string {
    for i := 0; i < steps; i++ {
		if this.Current.Prev == nil {
			break
		}
		this.Current = this.Current.Prev
	}
	return this.Current.Url
}


func (this *BrowserHistory) Forward(steps int) string {
    for i := 0; i < steps; i++ {
		if this.Current.Next == nil {
			break
		}
		this.Current = this.Current.Next
	}
	return this.Current.Url
}


/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := Constructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */