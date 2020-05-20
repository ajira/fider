package enum

type PostVisbility int

var (
	//EntityPrivate means that the post is visible only to users belonging
	//to same business entity as the user who created it and admin/moderators
	//EntityPrivate is the default post visibility
	EntityPrivate PostVisbility
	//Public means that the post is visible to everyone
	Public PostVisbility = 1
)
