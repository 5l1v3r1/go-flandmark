#include "flandmark_binding.h"
#include <flandmark_detector.h>

typedef struct {
  CvMemStorage * storage;
  CvSeq * rects;
} rects_list;

void * flandmark_binding_image_load(const char * filename) {
  return (void *)cvLoadImage(filename);
}

void * flandmark_binding_image_rgba(const uint8_t * rgba, int width,
                                    int height) {
  IplImage * img = cvCreateImageHeader(cvSize(width, height), IPL_DEPTH_8U, 4);
  cvSetData(img, buffer, img->widthStep);
  return (void *)img;
}

void * flandmark_binding_image_gray(const uint8_t * gray, int width,
                                    int height) {
  IplImage * img = cvCreateImageHeader(cvSize(width, height), IPL_DEPTH_8U, 1);
  cvSetData(img, buffer, img->widthStep);
  return (void *)img;
}

void flandmark_binding_image_free(void * image) {
  IplImage * frame = (IplImage *)image;
  cvReleaseImage(&frame);
}

void * flandmark_binding_cascade_load(const char * filename) {
  return (void *)cvLoad(filename, 0, 0, 0);
}

void * flandmark_binding_cascade_detect_objects(void * image, void * cascade,
  double factor, int minNeighbors, int minWidth, int minHeight, int maxWidth, int maxHeight) {
  IplImage * i = (IplImage *)image;
  CvHaarClassifierCascade * c = (CvHaarClassifierCascade *)cascade;
  CvMemStorage * storage = cvCreateMemStorage(0);
  cvClearMemStorage(storage);
  CvSeq * seq = cvHaarDetectObjects(i, c, storage, factor, minNeighbors,
    CV_HAAR_DO_CANNY_PRUNING, cvSize(minWidth, minHeight),
    cvSize(maxWidth, maxHeight));
  rects_list * res = new rects_list;
  res->storage = storage;
  res->rects = seq;
  return (void *)res;
}

void * flandmark_binding_cascade_free(void * cascade) {
  CvHaarClassifierCascade * c = (CvHaarClassifierCascade *)cascade;
  cvReleaseHaarClassifierCascade(&c);
}

int flandmark_binding_rects_count(void * rects) {
  rects_list * list = (rects_list *)rects;
  return list->rects->total;
}

void flandmark_binding_rects_get(void * rects, int idx, int * xywh) {
  rects_list * list = (rects_list *)rects;
  CvRect * r = (CvRect *)cvGetSeqElem(list->rects, idx);
  xywh[0] = r->x;
  xywh[1] = r->y;
  xywh[2] = r->width;
  xywh[3] = r->height;
}

void flandmark_binding_rects_free(void * rects) {
  rects_list * list = (rects_list *)rects;
  cvReleaseMemStorage(&list->storage);
  delete list;
}

void * flandmark_binding_model_init(const char * filename);
void * flandmark_binding_model_free(void * model);
uint8_t flandmark_binding_model_M(void * model);
int flandmark_binding_model_detect(void * model, void * image,
  double * landmarks, int * box, int * bwMargin);

