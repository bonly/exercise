package main 

import(
"errors"
)

type room struct{
	RoomID int `json:"RoomID"`;
	Address string;
	Pic string;
};

var rooms = []room{
	room{0, "珠江新城3316", "http://127.0.0.1/zj3316"},
	room{1, "琶洲新村1108", "http://127.0.0.1/pz1108"},
}

func getAllRoom() []room{
	return rooms;
}

func getRoomByID(id int)(*room, error){
	for _, aroom := range rooms{
		if aroom.RoomID == id{
			return &aroom, nil;
		}
	}
	return nil, errors.New("room not found");
}