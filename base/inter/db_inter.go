package inter

type DbInter interface {
	GetCon()
	Select()
	Del()
	Find()
	Up() int64
}
