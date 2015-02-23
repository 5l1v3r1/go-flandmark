#ifndef __FLANDMARK_BINDING_H__
#define __FLANDMARK_BINDING_H__

#include <stdint.h>
#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
  int x;
  int y;
  int width;
  int height;
} Rectangle;

typedef struct {
  int status;
  double * coords;
} DetectResult;

void * flandmark_binding_image_rgba(void * rgba, int width, int height);
void * flandmark_binding_image_gray(void * gray, int width, int height);
void flandmark_binding_image_free(void * image);

void * flandmark_binding_cascade_load(char * filename);
void * flandmark_binding_cascade_detect_objects(void * cascade, void * image,
  double factor, int minNeighbors, int minWidth, int minHeight, int maxWidth, int maxHeight);
void flandmark_binding_cascade_free(void * cascade);

int flandmark_binding_rects_count(void * rects);
Rectangle flandmark_binding_rects_get(void * rects, int idx);
void flandmark_binding_rects_free(void * rects);

void * flandmark_binding_model_init(char * filename);
void flandmark_binding_model_free(void * model);
uint8_t flandmark_binding_model_M(void * model);
DetectResult flandmark_binding_model_detect(void * model, void * image,
  Rectangle box);

#ifdef __cplusplus
}
#endif

#endif
