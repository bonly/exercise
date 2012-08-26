/**
 *  @file 20100509_bezier_curve.cpp
 *
 *  @date 2012-4-16
 *  @author Bonly
 */

typedef struct
{
        float x;
        float y;
} Point2D;

/* cp 在此是四个元素的数组:
 cp[0] 为起点，或上图中的 P0
 cp[1] 为第一控制点，或上图中的 P1
 cp[2] 为第二控制点，或上图中的 P2
 cp[3] 为结束点，或上图中的 P3
 t 为参数值，0 <= t <= 1 */
Point2D PointOnCubicBezier(Point2D* cp, float t)
{
    float ax, bx, cx;
    float ay, by, cy;
    float tSquared, tCubed;
    Point2D result;
    /* 计算多项式系数 */
    cx = 3.0 * (cp[1].x - cp[0].x);
    bx = 3.0 * (cp[2].x - cp[1].x) - cx;
    ax = cp[3].x - cp[0].x - cx - bx;
    cy = 3.0 * (cp[1].y - cp[0].y);
    by = 3.0 * (cp[2].y - cp[1].y) - cy;
    ay = cp[3].y - cp[0].y - cy - by;
    /* 计算t位置的点值 */
    tSquared = t * t;
    tCubed = tSquared * t;
    result.x = (ax * tCubed) + (bx * tSquared) + (cx * t) + cp[0].x;
    result.y = (ay * tCubed) + (by * tSquared) + (cy * t) + cp[0].y;
    return result;
}

/* ComputeBezier 以控制点 cp 所产生的曲线点，填入 Point2D 结构数组。
 调用方必须分配足够的空间以供输出，<sizeof(Point2D) numberOfPoints> */
void ComputeBezier(Point2D* cp, int numberOfPoints, Point2D* curve)
{
    float dt;
    int i;
    dt = 1.0 / (numberOfPoints - 1);
    for (i = 0; i < numberOfPoints; i++)
        curve[i] = PointOnCubicBezier(cp, i * dt);
}

/**
 *
 经典的曲线逼近方法，称作Bezier曲线。想必学过图形图像的都应该知道啦，所以概念性问题就不说啦。
该曲线分为一次/二次/三次/多次贝塞尔曲线，之所以这么分是为了更好的理解其中的内涵。

一次贝塞尔曲线，实际上就是一条连接两点的直线段。
二次贝塞尔曲线，就是两点间的一条抛物线，利用一个控制点来控制抛物线的形状。
三次贝塞尔曲线，则需要一个起点，一个终点，两个控制点来控制曲线的形状。
 */
