#include <iostream>
#include <cmath>
#include <vector>

#include <opencv2/highgui/highgui.hpp>
#include <opencv2/imgproc/imgproc.hpp>
#include <opencv2/imgproc/imgproc_c.h>

/* angle: finds a cosine of angle between vectors, from pt0->pt1 and from pt0->pt2
 */
double angle(cv::Point pt1, cv::Point pt2, cv::Point pt0)
{
    double dx1 = pt1.x - pt0.x;
    double dy1 = pt1.y - pt0.y;
    double dx2 = pt2.x - pt0.x;
    double dy2 = pt2.y - pt0.y;
    return (dx1*dx2 + dy1*dy2)/sqrt((dx1*dx1 + dy1*dy1)*(dx2*dx2 + dy2*dy2) + 1e-10);
}

/* findSquares: returns sequence of squares detected on the image
 */
void findSquares(const cv::Mat& src, std::vector<std::vector<cv::Point> >& squares)
{
    cv::Mat src_gray;
    cv::cvtColor(src, src_gray, cv::COLOR_BGR2GRAY);

    // Blur helps to decrease the amount of detected edges
    cv::Mat filtered;
    cv::blur(src_gray, filtered, cv::Size(3, 3));
    cv::imwrite("out_blur.jpg", filtered);

    // Detect edges
    cv::Mat edges;
    int thresh = 128;
    cv::Canny(filtered, edges, thresh, thresh*2, 3);
    cv::imwrite("out_edges.jpg", edges);

    // Dilate helps to connect nearby line segments
    cv::Mat dilated_edges;
    cv::dilate(edges, dilated_edges, cv::Mat(), cv::Point(-1, -1), 2, 1, 1); // default 3x3 kernel
    cv::imwrite("out_dilated.jpg", dilated_edges);

    // Find contours and store them in a list
    std::vector<std::vector<cv::Point> > contours;
    cv::findContours(dilated_edges, contours, cv::RETR_LIST, cv::CHAIN_APPROX_SIMPLE);

    // Test contours and assemble squares out of them
    std::vector<cv::Point> approx;
    for (size_t i = 0; i < contours.size(); i++)
    {
        // approximate contour with accuracy proportional to the contour perimeter
        cv::approxPolyDP(cv::Mat(contours[i]), approx, cv::arcLength(cv::Mat(contours[i]), true)*0.02, true);

        // Note: absolute value of an area is used because
        // area may be positive or negative - in accordance with the
        // contour orientation
        if (approx.size() == 4 && std::fabs(contourArea(cv::Mat(approx))) > 1000 &&
            cv::isContourConvex(cv::Mat(approx)))
        {
            double maxCosine = 0;
            for (int j = 2; j < 5; j++)
            {
                double cosine = std::fabs(angle(approx[j%4], approx[j-2], approx[j-1]));
                maxCosine = MAX(maxCosine, cosine);
            }

            if (maxCosine < 0.3)
                squares.push_back(approx);
        }
    }
}

/* findLargestSquare: find the largest square within a set of squares
 */
void findLargestSquare(const std::vector<std::vector<cv::Point> >& squares,
                       std::vector<cv::Point>& biggest_square)
{
    if (!squares.size())
    {
        std::cout << "findLargestSquare !!! No squares detect, nothing to do." << std::endl;
        //return;
    }

    int max_width = 0;
    int max_height = 0;
    int max_square_idx = 0;
    for (size_t i = 0; i < squares.size(); i++)
    {
        // Convert a set of 4 unordered Points into a meaningful cv::Rect structure.
        cv::Rect rectangle = cv::boundingRect(cv::Mat(squares[i]));

        //std::cout << "find_largest_square: #" << i << " rectangle x:" << rectangle.x << " y:" << rectangle.y << " " << rectangle.width << "x" << rectangle.height << endl;

        // Store the index position of the biggest square found
        if ((rectangle.width >= max_width) && (rectangle.height >= max_height))
        {
            max_width = rectangle.width;
            max_height = rectangle.height;
            max_square_idx = i;
        }
    }

    biggest_square = squares[max_square_idx];
}

int main(int argc, char *argv[])
{
    cv::Mat src = cv::imread(argv[1]);
    if (src.empty())
    {
        std::cout << "!!! Failed to open image" << std::endl;
        return -1;
    }

    std::vector<std::vector<cv::Point> > squares;
    findSquares(src, squares);

    // Draw all detected squares
    cv::Mat src_squares = src.clone();
    for (size_t i = 0; i < squares.size(); i++)
    {
        const cv::Point* p = &squares[i][0];
        int n = (int)squares[i].size();
        cv::polylines(src_squares, &p, &n, 1, true, cv::Scalar(0, 255, 0), 2, CV_AA);
    }
    cv::imwrite("out_squares.jpg", src_squares);
    cv::imshow("Squares", src_squares);

    std::vector<cv::Point> largest_square;
    findLargestSquare(squares, largest_square);

    // Draw circles at the corners
    for (size_t i = 0; i < largest_square.size(); i++ )
        cv::circle(src, largest_square[i], 4, cv::Scalar(0, 0, 255), cv::FILLED);
    cv::imwrite("out_corners.jpg", src);

    cv::imshow("Corners", src);
    cv::waitKey(0);

    return 0;
}
/*
g++ 20140509_square.cc -fpermissive -l opencv_core -l opencv_imgcodecs -l opencv_imgproc -lopencv_highgui
*/