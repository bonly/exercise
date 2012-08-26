/******************************************************************* 
 * Header file generated by Protoc for Embedded C.                 *
 * Version 0.2 (2012-01-31)                                        *
 *                                                                 *
 * Copyright (c) 2009-2012                                         *
 * Technische Universitaet Muenchen                                *
 * http://www4.in.tum.de/                                          *
 *                                                                 *
 * Source : 20100606_chat.proto
 * Package: 
 *                                                                 *
 * Do not edit.                                                    *
 *******************************************************************/

#include <20100603_chat.proto.h>

#define MAX_REPEATED_LEN 32
#define MAX_STRING_LEN 32
#define MAX_BYTES_LEN 32

/*******************************************************************
 * General functions
 *******************************************************************/

/*
 * returns the size of a length delimited message which also 
 * contains the first bytes for the length encoding.
 */
unsigned long Message_get_delimited_size(void *_buffer, int offset);

/*
 * Tests whether a message can be completely read from the given buffer at
 * the offset. The bytes [offset..offset+length-1] are interpreted.
 *
 * Returns 1 (true) if buffer[offset..offset+length-1] contains a complete
 * message or 0 (false) otherwise.
 */
int Message_can_read_delimited_from(void *_buffer, int offset, int length);


/*******************************************************************
 * Message: 20100606_chat.proto, line 3
 *******************************************************************/

/* Maximum size of a serialized ACTION-message, useful for buffer allocation. */
#define MAX_ACTION_SIZE 68

/* Structure that holds a deserialized ACTION-message. */
struct ACTION {
  struct ID _from_id;
  struct ID _to_id;
  signed long _action;
  signed long _kill;
  signed long _world;
  signed long _gold;
  signed long _attack;
};
/*
 * Serialize a ACTION-message into the given buffer at offset and return
 * new offset for optional next message.
 */
int ACTION_write_delimited_to(struct ACTION *_ACTION, void *_buffer, int offset);

/*
 * Serialize a ACTION-message together with its tag into the given buffer 
 * at offset and return new offset for optional next message.
 * Is useful if a ACTION-message is embedded in another message.
 */
int ACTION_write_with_tag(struct ACTION *_ACTION, void *_buffer, int offset, int tag);

/*
 * Deserialize a ACTION-message from the given buffer at offset and return
 * new offset for optional next message.
 *
 * Note: All fields in _ACTION will be reset to 0 before _buffer is interpreted.
 */
int ACTION_read_delimited_from(void *_buffer, struct ACTION *_ACTION, int offset);
