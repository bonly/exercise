/*
 *  @file Paladin.hpp
 *
 *  Created on: 2012-8-1
 *      Author: bonly
 */

#ifndef __PALADIN_HPP__
#define __PALADIN_HPP__
#include "20100627_pro.h"
//extern"C"{
#include "20100627_enum.h"
//}
enum {HEARTLEN=11};
enum
{
    HEARTBEAT = 0,
    DOWNLOAD_ARRAY,
    DOWNLOAD_ARRAY_ACK,
    UPLOAD_ARRAY,
    UPLOAD_ARRAY_ACK,
    CHALLENGE,
    CHALLENGE_ACK,
    PVP_LIST,
    PVP_LIST_ACK,
    MAX_PROTOCOL_CMD
};
CMD(Heartbeat);
#endif /* __PALADIN_HPP__*/
