#include "flandmark_binding.h"
#include "flandmark_detector.h"
#include <string.h>
#include <stdio.h>

typedef struct {
  CvMemStorage * storage;
  CvSeq * rects;
} rects_list;

void * flandmark_binding_image_rgba(void * rgba, int width, int height) {
  IplImage * img = cvCreateImage(cvSize(width, height), IPL_DEPTH_8U, 4);
  if (!img) {
    free(rgba);
    return NULL;
  }
  memcpy((void *)img->imageData, rgba, (size_t)width * (size_t)height * 4);
  free(rgba);
  return (void *)img;
}

void * flandmark_binding_image_gray(void * gray, int width, int height) {
  IplImage * img = cvCreateImage(cvSize(width, height), IPL_DEPTH_8U, 1);
  if (!img) {
    free(gray);
    return NULL;
  }
  memcpy((void *)img->imageData, gray, (size_t)width * (size_t)height);
  free(gray);
  return (void *)img;
}

void flandmark_binding_image_free(void * image) {
  IplImage * frame = (IplImage *)image;
  cvReleaseImage(&frame);
}

void * flandmark_binding_cascade_load(char * filename) {
  void * res = (void *)cvLoad(filename, 0, 0, 0);
  free(filename);
  return res;
}

void * flandmark_binding_cascade_detect_objects(void * cascade, void * image,
  double factor, int minNeighbors, int minWidth, int minHeight, int maxWidth,
  int maxHeight) {
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

void flandmark_binding_cascade_free(void * cascade) {
  CvHaarClassifierCascade * c = (CvHaarClassifierCascade *)cascade;
  cvReleaseHaarClassifierCascade(&c);
}

int flandmark_binding_rects_count(void * rects) {
  rects_list * list = (rects_list *)rects;
  return list->rects->total;
}

Rectangle flandmark_binding_rects_get(void * rects, int idx) {
  rects_list * list = (rects_list *)rects;
  CvRect * r = (CvRect *)cvGetSeqElem(list->rects, idx);
  
  Rectangle res = {r->x, r->y, r->width, r->height};
  return res;
}

void flandmark_binding_rects_free(void * rects) {
  rects_list * list = (rects_list *)rects;
  cvReleaseMemStorage(&list->storage);
  delete list;
}

void * flandmark_binding_model_init(char * filename) {
  void * res = (void *)flandmark_init(filename);
  free(filename);
  return res;
}

void flandmark_binding_model_free(void * model) {
  flandmark_free((FLANDMARK_Model *)model);
}

uint8_t flandmark_binding_model_M(void * model) {
  FLANDMARK_Model * m = (FLANDMARK_Model *)model;
  return m->data.options.M;
}

DetectResult flandmark_binding_model_detect(void * model, void * image,
  Rectangle box) {
  FLANDMARK_Model * m = (FLANDMARK_Model *)model;
  double * result = (double *)malloc(sizeof(double) * m->data.options.M * 2);
  int boxList[4] = {box.x, box.y, box.x+box.width, box.y+box.height};
  int numVal = flandmark_detect((IplImage *)image, boxList, m, result);
  if (numVal != 0) {
    free(result);
    DetectResult res = {numVal, NULL};
    return res;
  } else {
    DetectResult res = {0, result};
    return res;
  }
}

