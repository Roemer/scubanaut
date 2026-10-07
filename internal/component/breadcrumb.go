package component

type Breadcrumb struct {
	Items []BreadcrumbItem
}

func NewBreadcrumb(items ...BreadcrumbItem) Breadcrumb {
	lastIndex := len(items) - 1
	for i := range items {
		items[i].Active = i == lastIndex
	}
	return Breadcrumb{
		Items: items,
	}
}

type BreadcrumbItem struct {
	Name   string
	Link   string
	Active bool
}

func NewBreadcrumbItem(name, link string) BreadcrumbItem {
	return BreadcrumbItem{
		Name: name,
		Link: link,
	}
}
