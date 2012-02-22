/**
 *  @file b2LoopShape.h
 *
 *  @date 2012-2-21
 *  @Author: Bonly
 */

#ifndef B2LOOPSHAPE_H_
#define B2LOOPSHAPE_H_

#include "b2Shape.h"

class b2EdgeShape;

/// A loop shape is a free form sequence of line segments that form a circular list.
/// The loop may cross upon itself, but this is not recommended for smooth collision.
/// The loop has double sided collision, so you can use inside and outside collision.
/// Therefore, you may use any winding order.
class b2LoopShape : public b2Shape
{
public:
        b2LoopShape();

        /// Implement b2Shape.
        b2Shape* Clone(b2BlockAllocator* allocator) const;

        /// @see b2Shape::GetChildCount
        int32 GetChildCount() const;

        /// Get a child edge.
        void GetChildEdge(b2EdgeShape* edge, int32 index) const;

        /// This always return false.
        /// @see b2Shape::TestPoint
        bool TestPoint(const b2Transform& transform, const b2Vec2& p) const;

        /// Implement b2Shape.
        bool RayCast(b2RayCastOutput* output, const b2RayCastInput& input,
                                        const b2Transform& transform, int32 childIndex) const;

        /// @see b2Shape::ComputeAABB
        void ComputeAABB(b2AABB* aabb, const b2Transform& transform, int32 childIndex) const;

        /// Chains have zero mass.
        /// @see b2Shape::ComputeMass
        void ComputeMass(b2MassData* massData, float32 density) const;

        /// The vertices. These are not owned/freed by the loop shape.
        b2Vec2* m_vertices;

        /// The vertex count.
        int32 m_count;
};

inline b2LoopShape::b2LoopShape()
{
        m_type = e_loop;
        m_radius = b2_polygonRadius;
        m_vertices = NULL;
        m_count = 0;
}

#endif
