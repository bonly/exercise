#ifndef __GAMEOBJ_HPP_
#define __GAMEOBJ_HPP_

#include <utility>
#include <map>
#include <vector>
#include <boost/any.hpp>
#include <boost/shared_ptr.hpp>
#include <boost/make_shared.hpp>
#include <boost/enable_shared_from_this.hpp>

	
namespace b = boost;	
namespace bus{
namespace srv{

enum {PALADIN_IN_MAX_GROUP=3};
enum {MAX_PALADIN_IN_ARR=5};
#define GEST_CO_OPERATER 1/3
	
class Actor;
class Member;
class Player;
class Paladin;
class Paladin_attr;
class Level_attr;
class Addtion;
struct Comb;
struct Effects;
struct Addtion;
struct Buff;

typedef std::pair<int, b::any> item;
typedef std::map<int, b::any> items;
typedef items::iterator item_it;
typedef b::shared_ptr<Paladin_attr> paladin_attr_ptr;
typedef b::shared_ptr<Level_attr> paladin_level_ptr;
typedef b::shared_ptr<Paladin> paladin_ptr;
typedef b::shared_ptr<Paladin> player_member_type_ptr;

typedef std::pair<int, player_member_type_ptr> member_type;

/**
 * @brief 级别属性
 */
struct Level_attr
{
        int level; ///< 级别
        float lvl_minap; ///< 等级最少攻击
        float lvl_maxap; ///< 等级最大攻击
        float lvl_defence; ///< 等级防御
        float lvl_dr; ///< 等级伤害减免
        float lvl_hp; ///< 等级血量
        float lvl_aspeed; ///< 攻速
};

/**
 * @brief 影响关系
 */
struct Relation{
        Member* owner;   ///< 拥有者
        Member* driving; ///< 主动者
        Member* passive; ///< 被动者

        Relation():owner(0), driving(0), passive(0){}
};

/**
 * @brief 影响
 */
struct Effects{
        enum{EQUIPMENT=1,GROUP_SKILL,PERSON_SKILL}; ///type 的值
        Effects():times(1),type(0){}
        virtual int effects(Relation *)  {
            return 0;
        }
        virtual ~Effects() {
        }
        virtual int get_id(){};
        bool compare_by_id(int key) {
            return get_id() == key;
        }
        int times; ///< 影响次数,装备只会累一次属性
        int type;
};

/**
 * @brief 武侠属性
 */
class Paladin_attr
{
    public:
        Paladin_attr():
            paladin_id(0),
            valiant_class(0),
            attackco(.0),
            hplvlco(.0),
            defenceco (.0),
            aspeedco(.0),
            pic(0),
            templateid(0),
            skill_id(0),
            groups(PALADIN_IN_MAX_GROUP)
        {};
        int get_attr(const int template_id, const int valiant_class, b::shared_ptr<Paladin_attr> &pa);
        virtual ~Paladin_attr(){
            groups.clear();
        }

    public:
        signed int paladin_id; ///< 大侠ID(paladin_template->plaldin_template_id)
        signed int valiant_class; ///< 星级
        float attackco; ///< 攻击指数
        float hplvlco; ///< 生命指数
        float defenceco; ///< 防御指数
        float aspeedco; ///< 攻速
        signed int pic; ///< 头像ID
        signed int body_pic; ///< 身体
        signed int templateid; ///< 大侠模板ID(paladin_template->template_id)
        int skill_id; ///<转生技能ID
        std::vector<int> groups;///< 阵容组合
        b::shared_ptr<Effects>  skill;///< 转生技能
};

/**
 * @brief 队员
 */
class Member : public b::enable_shared_from_this<Member>{
    public:
        Member():stat(DEAD){}
        virtual ~Member(){
            buff.clear();
            equipment.clear();
            gests.clear();
            debuff.clear();
        };
        virtual int calc_attr()=0;
        virtual paladin_attr_ptr& get_attr()=0;
        virtual int calc_ap(Member &pa){};
        virtual int get_level()=0;
        virtual int effects(std::vector<b::shared_ptr<srv::Effects> > &){};
        int get_match(boost::shared_ptr<Member> &sec);
        static bool alive(boost::shared_ptr<Member> pa);


    public:
        Actor *player; ///< 归属玩家
        signed int member_id; ///< 队员ID(bug->stuff_id = paladin->paladin_id)
        signed int pic_id; ///< 头像ID
        signed int battle_num; ///< 排位 0-4 5-9
        signed int array_num; ///< 阵列位 0-4
        enum {
            DEAD, LIVE
        } stat; ///< 状态

        float now_HP; ///< 当前血量
        float now_AttackPower; ///< 当前攻击
        float now_Defence; ///< 当前防御
        float now_Aspeed; ///< 当前攻速

        float minap; ///< 最少攻击
        float maxap; ///< 最大攻击
        float gest; ///< 武功值

        int skill_id; ///< 技能ID
        std::vector<boost::shared_ptr<srv::Effects> >  buff;///< 增益技能
        std::vector<boost::shared_ptr<srv::Effects> >  equipment; ///< 装备
        std::vector<boost::shared_ptr<srv::Effects> >  gests; ///< 内外功
        std::vector<boost::shared_ptr<srv::Effects> >  debuff;///< 减益技能
        ///< 触发技能

};


/**
 * @brief 武侠
 */
class Paladin : public Member{
    public:
        enum{
            DRLvlCo = 100, ///< 伤害减免等级因子
            DRConstant = 500 ///<伤害减免常数
        };
        Paladin():Member(),bag_id(-1)
        {
            attr = b::make_shared<Paladin_attr>();
            level = b::make_shared<Level_attr>();
            stat = Member::DEAD;
        }
        virtual ~Paladin(){}
        virtual int calc_attr(){};  ///< 计算属性
        virtual paladin_attr_ptr& get_attr()
        {
            return attr;
        }
        virtual int get_level()
        {
            return level->level;
        }
        int calc_ap_range(Paladin&); ///< 攻击力范围
        int calc_hp(Paladin&);  ///< 血量
        int calc_defence(Paladin&); ///< 防御
        int calc_aspeed(Paladin &pa); ///< 攻速
        static int calc_ap(Paladin &pa);  ///< 计算当次攻击力
        int calc_injure();
        int calc_gest(Paladin &pa); ///< 计算武功值

    public:
        int bag_id; ///< 库表中的索引
        boost::shared_ptr<Paladin_attr> attr; ///< 属性
        boost::shared_ptr<Level_attr> level; ///< 级别属性
        ///< 装备
        ///< 套装容器
};

/**
 * @brief 玩家
 */
class Actor : public boost::enable_shared_from_this<Actor>
{
    public:
        Actor():player_id(0),
        lives(0),
        crew(0),
        summon(false),
        members(MAX_PALADIN_IN_ARR)
        {}
        virtual ~Actor(){
            members.clear();
        }
        virtual int convene()=0;
        virtual int group_buf()=0;
        virtual int call_buf(){return 0;}
        virtual int call_equip(){return 0;}
        virtual int call_debuf(){return 0;}
        virtual boost::shared_ptr<Member> get_live(){};
        virtual float team_blood()=0;
        virtual int take_equipment(){};

    public:
        int player_id; ///< ID
        int lives; ///< 存活成员数
        int crew; ///< 全体成员数
        bool summon; ///<是否为战斗发起者
        boost::shared_ptr<Actor> opponent; ///< 对手
        std::vector<boost::shared_ptr<Member> > members; ///< 队员
};

/**
 * @brief 真正的玩家对象
 */
class Player : public Actor
{
    public:
        virtual ~Player(){};
        virtual int convene();  ///< 召集大侠组成阵容
        virtual int group_buf(){};  ///< 发现组buf
        virtual int call_buf(){};
        virtual int call_equip(){};
        virtual int call_debuf(){};
        virtual float calc_gest(){};
        virtual int take_equipment(){};

        virtual float team_blood()  {
            float blood = 0;
            for (int i = 0; i < MAX_PALADIN_IN_ARR; ++i)      {
                if(Member::alive(members[i]))
                    blood += members[i]->now_HP;
            }
            return blood;
        }

        Player() : Actor()  {
            for (int i=0; i<MAX_PALADIN_IN_ARR; ++i)   {
                members[i] = boost::make_shared<Paladin>();
            }
        }

        void set_data(){};
        void set_base_attr(){};
};

/**
 * @brief 怪物
 */
class Monster : public Member{
    public:
    	  virtual ~Monster(){};
        virtual int calc_attr() {return 0;}
        virtual paladin_attr_ptr& get_attr() {
            return null_attr;
        }
        virtual int get_level() {
            return level;
        };
    public:
        paladin_level_ptr attr; ///< 属性
        paladin_attr_ptr null_attr;///< 没有使用的属性
        int icon_id;///< 身体图片
        int level;
};

/**
 * @brief 怪物组
 */
class MonsterTeam : public Actor
{
    public:
        virtual int convene();
        virtual int group_buf()
        {
            return 0;
        };
        virtual int call_buf()
        {
            return 0;
        };

        virtual float team_blood(){return 0.0;}

        void set_data();

        MonsterTeam() : Actor()
        {
            for (int i=0; i<MAX_PALADIN_IN_ARR; ++i)
            {
                members[i] = boost::make_shared<Monster>();
                members[i]->stat = Member::DEAD;
            }
        }
        virtual ~MonsterTeam(){}
};




/**
 * @brief 技能
 */
class Skill : public Effects
{
    public:
        virtual int effects(Relation *);
        int call_skill(Skill* sk,  Relation* me);

    private:
        int call_comb_skill(Skill* sk, Relation* me);
        int call_neutral_skill(Skill *sk, Relation* me);
        int call_buff_skill(Skill* sk, Relation* me);
        int call_debuff_skill(Skill* sk, Relation* me);
        int patch_attr_by_skill(Skill* sk, Member* tar);

        int call_boss_skill(Skill* sk, Relation* target);

        int skill_translation_op(int skillop);
        int modify(float &target, int op, int type, float value);

        int get_id()
        {
            return skill_ID;
        }

    public:
        int skill_ID;   ///< 技能ID
        int skill_carrier_type; ///< 技能载体类型
        int skill_type;  ///< 技能类型 0：中，1：增益，2：减益，3：触发型伤害技能
        int skill_target_type;  ///< 作用目标类型
        int skill_target_num;  ///< 作用最大目标数
        int skill_action_type1;  ///< 效果1类型
        int action1_param_type;  ///< 效果1参数类型
        float action1_param;  ///< 效果1参数
        int skill_action_type2; ///< 效果2类型
        int action2_param_type; ///< 效果2参数类型
        float action2_param; ///< 效果2参数
};

/**
 * @brief 状态
 */
class Buff : public Effects
{
    public:
        Buff()
        {
        }
        Buff(boost::shared_ptr<Skill> sk):skill(sk)
        {
        }
        int get_id()
        {
            if(skill==0)
                return -1;
            return skill->skill_ID;
        }
        int effects(Relation *ra)
        {
            if(skill == 0)
                return 0;
            return skill->effects(ra);
        }
        boost::shared_ptr<Skill>  skill; ///< 技能数据
        boost::shared_ptr<Relation> relation; ///< 技能相关对象

};

/**
 * @brief 装备属性
 */
struct Equipment_attr
{
        int equipment_template_id;
        int template_id;  ///< 类型模板ID
        int type; ///< 类型
        int star; ///< 星级
        int max_star; ///< 最大星级
        int top_level; ///< 最高等级
        int ap_base_min;
        int ap_base_max;
        int ap_add;
        int defence_base;
        int defence_add;
        int hp_base;
        int hp_add;
        float att_speed_base;
        float att_speed_add;
        int icon_id;
};

/**
 * @brief 装备
 */
struct Equipment : public Effects
{
        int bag_id;
        int equipment_id;
        int slot_num;
        int ap_max;
        int ap_min;
        int defence;
        int hp;
        float att_speed;
        int level;
        int is_active; ///< 0:在 1:不在
        boost::shared_ptr<Equipment_attr> attr;

        int effects(Relation *);
        int recalc();
        int get_id()
        {
            if (attr == 0)
                return -1;
            return attr->equipment_template_id;
        }
};

/**
 * @brief 对象匹配运算子
 */
struct key_extractor
{
    const int& operator()(const Level_attr& x)const
    {
        return x.level;
    }
    const int& operator()(const boost::shared_ptr<Level_attr>& x)const
    {
      return x->level;
    }
    const int& operator()(const boost::shared_ptr<Paladin_attr>& x)const
    {
      return x->paladin_id;
    }
};


/**
 * @brief 组合技能定义
 */
struct Comb
{
        int groupID;  ///< 组ID
        int buffID;  ///< 技能ID
        int paladin_count;  ///< 组内所需成员数
        int paladin1;  ///< 成员1
        int paladin2;  ///< 成员2
        int paladin3;  ///< 成员3
        int paladin4;  ///< 成员4
        int paladin5;  ///< 成员5
        int paladin6;
        int paladin7;
        int paladin8;
};

typedef struct _t_fight_power
{
        int ap_min; ///< 攻击力
        int ap_max; ///< 最大攻击力
        int defence; ///< 防守力
        int hp;    ///< 生命值
        float speed; ///< 速度
        int power;  ///< 武功值
        int group1_id;
        int group2_id;
        int group3_id;
        int group1_active;
        int group2_active;
        int group3_active;
        int skill_id;
        int ret;  ///< 返回执行结果 -1:失败 0成功
        _t_fight_power():
            ap_min(0),ap_max(0),defence(0),hp(0),speed(0),
            power(0),group1_id(0),group2_id(1),group3_id(0),
            group1_active(0), group2_active(0), group3_active(0),
            skill_id(0),ret(1)
        {}
}TFightPower;

TFightPower CalPaladinPower(int player_id, int idx, int in_battle, int up_to_class, int up_to_level);


typedef boost::shared_ptr<Level_attr> level_type;
typedef boost::shared_ptr<Paladin_attr> paladin_type;

typedef boost::shared_ptr<srv::Skill> skill_ptr;
typedef std::map<int, boost::shared_ptr<srv::Skill> > skill_data;
typedef std::map<int, boost::shared_ptr<srv::Skill> >::iterator skill_data_map_it;

typedef boost::shared_ptr<srv::Comb> comb_ptr;
typedef std::map<int, boost::shared_ptr<srv::Comb> > comb_data;
typedef std::map<int, boost::shared_ptr<srv::Comb> >::iterator comb_data_map_it;
class ObjSrv{
	public:
	  int get_Paladin(const int player_id, const int bag_id, boost::shared_ptr<Paladin> &pa, int &array_num);
	  int get_Paladin(const int player_id,
            std::vector<boost::shared_ptr<Member> > &paladin,
            int &paladin_count, const int count);
	  	
	public:
		skill_data _skill_data;
		comb_data _comb_data;
		std::map<int, level_type> _lvl;
		std::map<int, boost::shared_ptr<Equipment_attr> > _equipment_attr;
		std::map<int, paladin_type> _paladin;
};

}
}

#endif
