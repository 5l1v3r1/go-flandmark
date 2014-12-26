#ifndef __FLANDMARK_BINDING_H__
#define __FLANDMARK_BINDING_H__

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

void * flandmark_binding_image_rgba(const uint8_t * rgba, int width,
  int height);
void * flandmark_binding_image_gray(const uint8_t * gray, int width,
  int height);
void flandmark_binding_image_free(void * image);

void * flandmark_binding_cascade_load(const char * filename);
void * flandmark_binding_cascade_detect_objects(void * cascade, void * image,
  double factor, int minNeighbors, int minWidth, int minHeight, int maxWidth, int maxHeight);
void flandmark_binding_cascade_free(void * cascade);

int flandmark_binding_rects_count(void * rects);
void flandmark_binding_rects_get(void * rects, int idx, int * xywh);
void flandmark_binding_rects_free(void * rects);

void * flandmark_binding_model_init(const char * filename);
void flandmark_binding_model_free(void * model);
uint8_t flandmark_binding_model_M(void * model);
int flandmark_binding_model_detect(void * model, void * image,
  double * landmarks, int * box);

#ifdef __cplusplus
}
#endif

#endif
