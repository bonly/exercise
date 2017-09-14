package main 

import (
"log"
"fmt"
"time"
"sort"
)

type Room struct{
	Hotel_id string;
	Room_id string;
	Last_time time.Time;
	Last_faile bool;
};

// func (rm Room) Less(i Room, j Room) bool{
// 	return i.Hotel_id < j.Hotel_id && i.Room_id < j.Room_id;
// }

type Rooms []Room;

func (ls Rooms) Len() int{
	return len(ls);
}

func (ls Rooms) Swap(i int, j int){
	ls[i], ls[j] = ls[j], ls[i];
}

func (ls Rooms) Less(i int, j int) bool{
	return ls[i].Room_id + ls[i].Hotel_id < ls[j].Room_id + ls[j].Hotel_id;
}

var all_room Rooms;

func get_local_room_state(hotel_id string, room_id string, tn time.Time) *Room{
	if !sort.IsSorted(all_room){
		sort.Sort(all_room);
	}
	aroom := Room{
		Hotel_id : hotel_id,
		Room_id : room_id,
		Last_time : tn,
		Last_faile : true,
	};
	idx := sort.Search(len(all_room), func(i int) bool{
		// fmt.Printf("查 %v %v\n", all_room[i].Hotel_id, all_room[i].Room_id);
		return  all_room[i].Room_id + all_room[i].Hotel_id >= aroom.Room_id + aroom.Hotel_id;
	});	
	if idx < len(all_room) && all_room[idx].Room_id + all_room[idx].Hotel_id == aroom.Room_id + aroom.Hotel_id { //找到
		fmt.Println("找到");
		return &all_room[idx];
	}else{ //找不到，新增
		fmt.Println("找不到");
		all_room = append(all_room, aroom);
		return &all_room[idx];
	}
}

func main(){
	tn := time.Now();
	all_room = append(all_room, Room{"1","6",tn,true}, 
								Room{"2","1",tn,false}, 
								Room{"1","3",tn,false},
								Room{"1","5",tn,false},
	);


	// fmt.Printf("all room: \n %+v \n", all_room);
	room := get_local_room_state("1", "5", tn);
	log.Println(fmt.Sprintf("%v", room));
	for idx, pt := range all_room{
		fmt.Printf("search add 后数据: [%d] %v\n", idx, pt);
	}

	room = get_local_room_state("1", "1", tn);
	log.Println(fmt.Sprintf("%v", room));
	for idx, pt := range all_room{
		fmt.Printf("search add 数据: [%d] %v\n", idx, pt);
	}	
}