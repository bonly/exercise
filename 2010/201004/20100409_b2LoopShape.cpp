/**
 *  @file b2LoopShape.cpp
 *
 *  @date 2012-2-21
 *  @Author: Bonly
 */

#include "b2LoopShape.h"
#include <Box2D/Collision/Shapes/b2LoopShape.h>
#include <Box2D/Collision/Shapes/b2EdgeShape.h>
#include <new>

b2Shape* b2LoopShape::Clone(b2BlockAllocator* allocator) const
{
        void* mem = allocator->Allocate(sizeof(b2LoopShape));
        b2LoopShape* clone = new (mem) b2LoopShape;
        *clone = *this;
        clone->m_vertices = (b2Vec2*)allocator->Allocate(m_count * sizeof(b2Vec2));
        return clone;
}

int32 b2LoopShape::GetChildCount() const
{
        return m_count;
}

void b2LoopShape::GetChildEdge(b2EdgeShape* edge, int32 index) const
{
        b2Assert(2 <= m_count);
        b2Assert(0 <= index && index < m_count);
        edge->m_type = b2Shape::e_edge;
        edge->m_radius = m_radius;
        edge->m_hasVertex0 = true;
        edge->m_hasVertex3 = true;

        int32 i0 = index - 1 >= 0 ? index - 1 : m_count - 1;
        int32 i1 = index;
        int32 i2 = index + 1 < m_count ? index + 1 : 0;
        int32 i3 = index + 2;
        while (i3 >= m_count)
        {
                i3 -= m_count;
        }

        edge->m_vertex0 = m_vertices[i0];
        edge->m_vertex1 = m_vertices[i1];
        edge->m_vertex2 = m_vertices[i2];
        edge->m_vertex3 = m_vertices[i3];

        edge->m_index1 = i1;
        edge->m_index2 = i2;
}

bool b2LoopShape::TestPoint(const b2Transform& xf, const b2Vec2& p) const
{
        B2_NOT_USED(xf);
        B2_NOT_USED(p);
        return false;
}

bool b2LoopShape::RayCast(b2RayCastOutput* output, const b2RayCastInput& input,
                                                        const b2Transform& xf, int32 childIndex) const
{
        b2Assert(childIndex < m_count);

        b2EdgeShape edgeShape;

        int32 i1 = childIndex;
        int32 i2 = childIndex + 1;
        if (i2 == m_count)
        {
                i2 = 0;
        }

        edgeShape.m_vertex1 = m_vertices[i1];
        edgeShape.m_vertex2 = m_vertices[i2];

        return edgeShape.RayCast(output, input, xf, 0);
}

void b2LoopShape::ComputeAABB(b2AABB* aabb, const b2Transform& xf, int32 childIndex) const
{
        b2Assert(childIndex < m_count);

        int32 i1 = childIndex;
        int32 i2 = childIndex + 1;
        if (i2 == m_count)
        {
                i2 = 0;
        }

        b2Vec2 v1 = b2Mul(xf, m_vertices[i1]);
        b2Vec2 v2 = b2Mul(xf, m_vertices[i2]);

        aabb->lowerBound = b2Min(v1, v2);
        aabb->upperBound = b2Max(v1, v2);
}

void b2LoopShape::ComputeMass(b2MassData* massData, float32 density) const
{
        B2_NOT_USED(density);

        massData->mass = 0.0f;
        massData->center.SetZero();
        massData->I = 0.0f;
}
