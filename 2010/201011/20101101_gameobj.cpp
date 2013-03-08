/**
 * @addtogroup GameObj
 * @file GameObj.cpp
 * @brief 游戏中存在的对象类定义
 *
 * @author bonly
 * @date 2012-8-31 bonly create
 * @date 2012-9-25 文档化
 */

#include "20101101_gameobj.hpp"
#include "20101102_service_pool.hpp"
#include <cmath>
#include <boost/bind.hpp>
#define VLOG(x) std::clog
	
namespace bus{
boost::shared_ptr<bus::ServicePool> g_spool;		
namespace srv{

/** @brief 预计算一个武将升级后的数据
 * @param player_id: 玩家ID
 * @param idx: 武将时指bag_paladin/paladin表中的idx,
 * @param in_battle: 0/1是在阵容中, 表示去查paladin,
 * @param up_to_class: 要转到的星级(转生用）
 * @param up_to_level: 要升到的级别
 **/
TFightPower CalPaladinPower(int player_id, int idx, int in_battle, int up_to_class, int up_to_level)
{
    TFightPower fp;
    fp.ret = -1;

    Player player;
    player.player_id = player_id;
    int array_num = idx;
    boost::shared_ptr<Paladin> pa;
    /// 取武侠数据
    if (in_battle == 1) ///- 阵容中的武侠
    {
        if (array_num > 4 || array_num < 0)
        {
            VLOG(5) << "阵位不正确，需(0-4). @array_num=" << array_num;
            return fp;
        }

        /*
        if (player.convene() < 0 || player.members[array_num]->stat != Member::LIVE)
        {
            VLOG(5) << "阵容中找到不array_num[" << array_num << "]的武侠";
            return fp;
        }
        */

        
        //player.group_buf(); /// {\n 阵容加成
        pa = b::dynamic_pointer_cast<Paladin>(player.members[array_num]);

        /// 在阵容中的武侠,取装备 \n}
        //bus::g_spool->_obj_srv->get_paladin_equipment(pa);
        
        //player.call_equip(); ///< 此处为计算所有阵容中武侠的装备
        //player.call_buf();  ///< 处理buf
        //player.call_debuf();

        fp.group1_id = pa->attr->groups[0];
        fp.group2_id = pa->attr->groups[1];
        fp.group3_id = pa->attr->groups[2];
//        VLOG(200) << "group skill: " << fp.group1_id;
//        VLOG(200) << "group skill: " << fp.group2_id;
//        VLOG(200) << "group skill: " << fp.group3_id;        
        for (unsigned int i=0; i<pa->attr->groups.size(); ++i)
        {
            if (pa->attr->groups[i] > 0 &&
                        (pa->buff.end() != find_if(pa->buff.begin(), pa->buff.end(),
                               b::bind(&Effects::get_id,_1) == pa->attr->groups[i])))
            {
                if (i == 0) {
                    fp.group1_active = 1;
                    VLOG(200) << "active group 1 for @paladin_id=" << pa->attr->paladin_id;
                }
                if (i == 1) {
                    fp.group2_active = 1;
                    VLOG(200) << "active group 2 for @paladin_id=" << pa->attr->paladin_id;
                }
                if (i == 2) {
                    fp.group3_active = 1;
                    VLOG(200) << "active group 3 for @paladin_id=" << pa->attr->paladin_id;
                }
            }
        }
    }
    else  ///- 阵容外的武侠从背包中指定的大侠
    {
        if(0 != bus::g_spool->_obj_srv->get_Paladin(player_id, idx, pa, array_num) || pa == 0)
        {
            VLOG(5) << "背包中找不到指定的武侠. @player_id=" << player_id << " @bag_id=" << idx;
            return fp;
        }
    }
    
    if(!bus::g_spool.get()){
    	  VLOG(5) << "err spool";
    }
    if(!bus::g_spool.get() || !(bus::g_spool->_obj_srv.get())){
	    	VLOG(5) << "error obj srv";
	    	return fp;
    }
    
    /// 修改级别
    pa->level = bus::g_spool->_obj_srv->_lvl[up_to_level];
    if(pa->level == 0)
    {
        VLOG(5) << "error level: " << up_to_level;
        return fp;
    }    

    /*
    /// 修改星级 @todo 重新挂个人技能
    b::shared_ptr<Paladin_attr> new_attr;
    if(0 != pa->attr->get_attr(pa->attr->templateid, up_to_class, new_attr))
    {
        VLOG(5) << "无此星级. @class=" << up_to_class;
        return fp;
    }
   
    pa->attr = new_attr;

    /// 计算基础属性
    pa->calc_attr();

    ///  装备加成
    pa->effects(pa->equipment);  ///计算此武侠的装备加成

    ///  计算buff
    pa->effects(pa->buff);

    /// 给返回值
    fp.ap_max = pa->maxap;
    fp.ap_min = pa->minap;
    fp.defence = pa->now_Defence;
    fp.hp = pa->now_HP;
    fp.speed = pa->now_Aspeed;
    fp.skill_id = pa->attr->skill_id;

    if (fp.speed == 0)
    {
        VLOG(5) << "error paladin data: att_speed 0";
        return fp;
    }
    */
    /// 武功值
    double tmp_power = (fp.ap_max + fp.ap_min)/2 * (double)fp.defence * (double)fp.hp / fp.speed;
    fp.power = (int)pow((double)tmp_power, (double)GEST_CO_OPERATER);
    fp.ret = 0;

    return fp;
}
/**
 * 召集大侠组成阵容,初始化基本属性
 * @return
 */
int Player::convene()
{
    //boost::shared_ptr<Member> pal[MAX_PALADIN_IN_ARR] = {members[0],members[1],members[2],members[3],members[4]};
    if (g_spool->_obj_srv->get_Paladin(player_id, members, lives, MAX_PALADIN_IN_ARR) != 0)
    {
        VLOG(5) << "角色Paladin数据错误. @palyer_id=" << player_id;
        return -1;
    }
    if (lives == 0 )
    {
        VLOG(5) << "角色没有Paladin阵容. @player_id=" << player_id;
        return -1;
    }
    crew = lives;
    //set_data();
    //set_base_attr();

    /* for debug
    for (int i=0; i<MAX_PALADIN_IN_ARR; ++i)
    {
        if (pal[i]->stat == Member::LIVE)
        {
            Paladin* pa= (Paladin*)pal[i].get();
            VLOG(200) << "paladin_id: " << pa->attr->paladin_id;
            VLOG(200) << "level: " << pa->level->level;
        }
    }
    //*/
    return 0;
}

/**
 取所有上阵武侠
 */
int ObjSrv::get_Paladin(const int player_id,
            std::vector<boost::shared_ptr<Member> > &paladin,
            int &paladin_count, const int count)
{
    //boost::unique_lock<boost::shared_mutex> lock(_mutex);
    /*
    if (0 != get_db())
    {
        VLOG(200) << "get db failed";
        return -1;
    }

    sprintf(_sql,
                " select p.paladin_template_id, p.level, p.array_num, p.bag_id, "  // 0-3
                "        v.attack_co, v.hp_lvl_co, v.defence_co, v.att_speed_co, " // 4-7
                "        m.min_ap, m.max_ap, m.defence, m.defence, m.hp, p.head_icon_id, "  // 8-13
                "        v.group1_id, v.group2_id, v.group3_id, v.template_id, p.paladin_id, v.body_icon_id " // 14-19
                " from player_bag_paladin p, paladin_template v, paladin_level_base m "
                " where stuff_type=0 and is_active=1 and array_num in(0,1,2,3,4) "
                "       and p.paladin_template_id = v.paladin_template_id and m.level = p.level "
                "       and p.player_id=%d "
                " order by array_num limit %d", player_id, count
                );
    if (0 != mysql_query(&_connect, _sql))
    {
        VLOG(5) << _sql;
        VLOG(5) << "query sql exception: " << mysql_error(&_connect);
        db_close();
        return -2;
    }
    if (NULL == (_result = mysql_store_result(&_connect)))
    {
        VLOG(5) << "query result exception: " << mysql_error(&_connect);
        mysql_free_result(_result);
        return -3;
    }
    paladin_count = mysql_num_rows(_result);
    int rows = 0;
    for ( ; NULL != (_row = mysql_fetch_row(_result)); ++rows){
    */
    for (int i=0; i<5; ++i){
        int array_num = i;
        if (array_num >= MAX_PALADIN_IN_ARR)
            continue;

        boost::shared_ptr<Paladin> pa = boost::dynamic_pointer_cast<Paladin>(paladin[array_num]);
        pa->member_id = i; ///bag->stuff_id=paladin->paladin_id

        ///检查数据是否正确
        int tmp = i;
        if (bus::g_spool->_obj_srv->_paladin.end() ==
                    g_spool->_obj_srv->_paladin.find(tmp)) {
            VLOG(5) << "no paladin: " << tmp;
            continue;
        }
        pa->attr = g_spool->_obj_srv->_paladin[tmp];

        tmp = i;
        if (bus::g_spool->_obj_srv->_lvl.end() ==
                    g_spool->_obj_srv->_lvl.find(tmp)) {
            VLOG(5) << "no paladin level: " << tmp;
            continue;
        }
        pa->level = g_spool->_obj_srv->_lvl[tmp];
        pa->stat = Member::LIVE;

        if (pa->attr->skill_id == 0)
            continue;
        ///检查skill是否存在,配置数据错误, @todo 考虑直接用内存只构造好的attr
        if (bus::g_spool->_obj_srv->_skill_data.end() ==
                    bus::g_spool->_obj_srv->_skill_data.find(pa->attr->skill_id)) {
            VLOG(5) << "skill data error: " << pa->attr->skill_id ;
            continue;
        } else {
            boost::shared_ptr<Buff> buf = b::make_shared<Buff>();
            buf->skill = bus::g_spool->_obj_srv->_skill_data[pa->attr->skill_id];
            buf->type = Effects::PERSON_SKILL;
            pa->buff.push_back(buf);
        }
    }
    //mysql_free_result(_result);
    return 0;
}

/**
 从背包中取武侠
 */
int ObjSrv::get_Paladin(const int player_id, const int bag_id, boost::shared_ptr<Paladin> &pa, int &array_num){
	
}

/**
 * @brief 检查大侠是否存活
 */
bool Member::alive(boost::shared_ptr<Member> pa)
{
    return (pa->stat == Member::LIVE) ? true : false;
}

}
}

