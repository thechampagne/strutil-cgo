// TODOs: func MapLines(str string, fn func(string) string) string
//        func SplitAndMap(str string, split string, fn func(string) string) string
package main

/*
#include <stdlib.h>

typedef enum {
  CENTER,
  LEFT,
  RIGHT,
} align_t;

typedef struct {
  char* top;
  char* top_right;
  char* right;
  char* bottom_right;
  char* bottom;
  char* bottom_left;
  char* left;
  char* top_left;
} box9_slice_t;

typedef struct {
  char* buffer;
  char* error;
} draw_box_t;

typedef struct {
  char* buffer;
  char* error;
} random_t;

typedef struct {
  char* buffer;
  char* error;
} sub_string_t;
*/
import "C"
import (
	"unsafe"
	"github.com/ozgio/strutil"
)

//export strutil_align
func strutil_align(str *C.char, align_to C.align_t, width C.int) *C.char {
  if align_to == C.CENTER {
    return C.CString(strutil.Align(C.GoString(str), strutil.Center, int(width)))
  }
  if align_to == C.LEFT {
    return C.CString(strutil.Align(C.GoString(str), strutil.Left, int(width)))
  }
  if align_to == C.RIGHT {
    return C.CString(strutil.Align(C.GoString(str), strutil.Right, int(width)))
  }
  return C.CString(strutil.Align(C.GoString(str), strutil.Center, int(width)))
}

//export strutil_align_center
func strutil_align_center(str *C.char, width C.int) *C.char {
  return C.CString(strutil.AlignCenter(C.GoString(str), int(width)))
}

//export strutil_align_left
func strutil_align_left(str *C.char) *C.char {
  return C.CString(strutil.AlignLeft(C.GoString(str)))
}

//export strutil_align_right
func strutil_align_right(str *C.char, width C.int) *C.char {
  return C.CString(strutil.AlignRight(C.GoString(str), int(width)))
}

//export strutil_center_text
func strutil_center_text(str *C.char, width C.int) *C.char {
  return C.CString(strutil.CenterText(C.GoString(str), int(width)))
}

//export  strutil_count_words
func strutil_count_words(str *C.char) C.int {
  return C.int(strutil.CountWords(C.GoString(str)))
}

//export strutil_draw_box
func strutil_draw_box(content *C.char, width C.int, align C.align_t) *C.draw_box_t {
  self := (*C.draw_box_t) (C.malloc(C.size_t(unsafe.Sizeof(C.draw_box_t{}))))
  al := strutil.Center
  if align == C.LEFT {
    al = strutil.Left
  } else if align == C.RIGHT {
    al = strutil.Right
  }
  res, err := strutil.DrawBox(C.GoString(content), int(width), al)
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
    self.buffer = C.CString(res)
    self.error = nil
    return self
}

//export strutil_draw_custom_box
func strutil_draw_custom_box(content *C.char, width C.int, align C.align_t, chars *C.box9_slice_t, str_new_line *C.char) *C.draw_box_t {
  self := (*C.draw_box_t) (C.malloc(C.size_t(unsafe.Sizeof(C.draw_box_t{}))))
  al := strutil.Center
  if align == C.LEFT {
    al = strutil.Left
  } else if align == C.RIGHT {
    al = strutil.Right
  }
  box := strutil.Box9Slice{
    Top: C.GoString(chars.top),
    TopRight: C.GoString(chars.top_right),
    TopLeft: C.GoString(chars.top_left),
    Right: C.GoString(chars.right),
    Left: C.GoString(chars.left),
    Bottom: C.GoString(chars.bottom),
    BottomRight: C.GoString(chars.bottom_right),
    BottomLeft: C.GoString(chars.bottom_left),
  }
  res, err := strutil.DrawCustomBox(C.GoString(content), int(width), al, box, C.GoString(str_new_line))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
    self.buffer = C.CString(res)
    self.error = nil
    return self
}

//export strutil_expand_tabs
func strutil_expand_tabs(str *C.char, count C.int) *C.char {
  return C.CString(strutil.ExpandTabs(C.GoString(str), int(count)))
}

//export strutil_indent
func strutil_indent(str *C.char, left *C.char) *C.char {
  return C.CString(strutil.Indent(C.GoString(str), C.GoString(left)))
}

//export strutil_is_ascii
func strutil_is_ascii(s *C.char) C.int {
  res := strutil.IsASCII(C.GoString(s))
  if res {
    return C.int(1)
  }
  return C.int(0)
}

// TODO: func MapLines(str string, fn func(string) string) string

//export strutil_must_sub_string
func strutil_must_sub_string(str *C.char, start C.int, end C.int) *C.char {
  return C.CString(strutil.MustSubstring(C.GoString(str), int(start), int(end)))
}
  
//export strutil_os_new_line
func strutil_os_new_line() *C.char {
  return C.CString(strutil.OSNewLine())
}

//export strutil_pad
func strutil_pad(str *C.char, width C.int, left_pad *C.char, right_pad *C.char) *C.char {
  return C.CString(strutil.Pad(C.GoString(str), int(width), C.GoString(left_pad), C.GoString(right_pad)))
}

//export strutil_pad_left
func strutil_pad_left(str *C.char, width C.int, pad *C.char) *C.char {
  return C.CString(strutil.PadLeft(C.GoString(str), int(width), C.GoString(pad)))
}

//export strutil_pad_right
func strutil_pad_right(str *C.char, width C.int, pad *C.char) *C.char {
  return C.CString(strutil.PadRight(C.GoString(str), int(width), C.GoString(pad)))
}

//export strutil_random
func strutil_random(str_set *C.char, length C.int) *C.random_t {
  self := (*C.random_t) (C.malloc(C.size_t(unsafe.Sizeof(C.random_t{}))))
  res, err := strutil.Random(C.GoString(str_set), int(length))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
    self.buffer = C.CString(res)
    self.error = nil
    return self
}

//export strutil_remove_accents
func strutil_remove_accents(str *C.char) *C.char {
  return C.CString(strutil.RemoveAccents(C.GoString(str)))
}

//export strutil_replace_all_to_one
func strutil_replace_all_to_one(str *C.char, from **C.char, from_length C.int , to *C.char) *C.char {
  var fromSlice []string
  slice := (*[1 << 30]*C.char)(unsafe.Pointer(from))[:from_length:from_length]
  for _, v := range slice {
    fromSlice = append(fromSlice, C.GoString(v))
  }
  return C.CString(strutil.ReplaceAllToOne(C.GoString(str), fromSlice, C.GoString(to)))
}

//export strutil_reverse
func strutil_reverse(s *C.char) *C.char {
  return C.CString(strutil.Reverse(C.GoString(s)))
}

//export strutil_slugify
func strutil_slugify(str *C.char) *C.char {
  return C.CString(strutil.Slugify(C.GoString(str)))
}

//export strutil_slugify_special
func strutil_slugify_special(str *C.char, delimiter *C.char) *C.char {
  return C.CString(strutil.SlugifySpecial(C.GoString(str), C.GoString(delimiter)))
}

//export strutil_splice
func strutil_splice(str *C.char, new_str *C.char, start C.int, end C.int) *C.char {
  return C.CString(strutil.Splice(C.GoString(str), C.GoString(new_str), int(start), int(end)))
}

// TODO: func SplitAndMap(str string, split string, fn func(string) string) string

//export strutil_split_camel_case
func strutil_split_camel_case(str *C.char) **C.char {
  res := strutil.SplitCamelCase(C.GoString(str))
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  slice := (*[1<<30 - 1]*C.char)(array)
  for i, v := range res {
    slice[i] = C.CString(v)
  }
  return (**C.char) (array)
}

//export strutil_sub_string
func strutil_sub_string(str *C.char, start C.int, end C.int) *C.sub_string_t {
  self := (*C.sub_string_t) (C.malloc(C.size_t(unsafe.Sizeof(C.sub_string_t{}))))
  res, err := strutil.Substring(C.GoString(str), int(start), int(end))
  if err != nil {
    self.buffer = nil
    self.error = C.CString(err.Error())
    return self
  }
    self.buffer = C.CString(res)
    self.error = nil
    return self
}

//export strutil_summary
func strutil_summary(str *C.char, length C.int, end *C.char) *C.char {
  return C.CString(strutil.Summary(C.GoString(str), int(length), C.GoString(end)))
}

//export strutil_tile
func strutil_tile(pattern *C.char, length C.int) *C.char {
  return C.CString(strutil.Tile(C.GoString(pattern), int(length)))
}

//export strutil_to_camel_case
func strutil_to_camel_case(str *C.char) *C.char {
  return C.CString(strutil.ToCamelCase(C.GoString(str)))
}

//export strutil_to_snake_case
func strutil_to_snake_case(str *C.char) *C.char {
  return C.CString(strutil.ToSnakeCase(C.GoString(str)))
}

//export strutil_word_wrap
func strutil_word_wrap(str *C.char, colLen C.int, break_long_words C.int) *C.char {
  if break_long_words != 1 {
    return C.CString(strutil.WordWrap(C.GoString(str), int(colLen), false))
  }
  return C.CString(strutil.WordWrap(C.GoString(str), int(colLen), true))
}

//export strutil_words
func strutil_words(str *C.char) **C.char {
  res := strutil.Words(C.GoString(str))
  array := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))
  slice := (*[1<<30 - 1]*C.char)(array)
  for i, v := range res {
    slice[i] = C.CString(v)
  }
  return (**C.char) (array)
}

//export strutil_default_box9_slice
func strutil_default_box9_slice() *C.box9_slice_t {
  self := (*C.box9_slice_t) (C.malloc(C.size_t(unsafe.Sizeof(C.box9_slice_t{}))))
  res := strutil.DefaultBox9Slice()
  self.top = C.CString(res.Top)
  self.top_right = C.CString(res.TopRight)
  self.top_left = C.CString(res.TopLeft)
  self.right = C.CString(res.Right)
  self.left = C.CString(res.Left)
  self.bottom = C.CString(res.Bottom)
  self.bottom_right = C.CString(res.BottomRight)
  self.bottom_left = C.CString(res.BottomLeft)
  return self
}

//export strutil_simple_box9_slice
func strutil_simple_box9_slice() *C.box9_slice_t {
  self := (*C.box9_slice_t) (C.malloc(C.size_t(unsafe.Sizeof(C.box9_slice_t{}))))
  res := strutil.SimpleBox9Slice()
  self.top = C.CString(res.Top)
  self.top_right = C.CString(res.TopRight)
  self.top_left = C.CString(res.TopLeft)
  self.right = C.CString(res.Right)
  self.left = C.CString(res.Left)
  self.bottom = C.CString(res.Bottom)
  self.bottom_right = C.CString(res.BottomRight)
  self.bottom_left = C.CString(res.BottomLeft)
  return self
}

//export box9_slice_clean
func box9_slice_clean(self *C.box9_slice_t) {
  if self != nil {
    if self.top != nil { C.free(unsafe.Pointer(self.top)) }
    if self.top_right != nil { C.free(unsafe.Pointer(self.top_right)) }
    if self.right != nil { C.free(unsafe.Pointer(self.right)) }
    if self.bottom_right != nil { C.free(unsafe.Pointer(self.bottom_right)) }
    if self.bottom != nil { C.free(unsafe.Pointer(self.bottom)) }
    if self.bottom_left != nil { C.free(unsafe.Pointer(self.bottom_left )) }
    if self.left != nil { C.free(unsafe.Pointer(self.left)) }
    if self.top_left != nil { C.free(unsafe.Pointer(self.top_left)) }
    C.free(unsafe.Pointer(self))
  }
}

//export draw_box_clean
func draw_box_clean(self *C.draw_box_t) {
  if self != nil  {
    if self.buffer != nil { C.free(unsafe.Pointer(self.buffer)) }
    if self.error != nil { C.free(unsafe.Pointer(self.error)) }
    C.free(unsafe.Pointer(self))
  }
}

//export random_clean
func random_clean(self *C.random_t) {
  if self != nil  {
    if self.buffer != nil { C.free(unsafe.Pointer(self.buffer)) }
    if self.error != nil { C.free(unsafe.Pointer(self.error)) }
    C.free(unsafe.Pointer(self))
  }
}

//export sub_string_clean
func sub_string_clean(self *C.sub_string_t) {
  if self != nil  {
    if self.buffer != nil { C.free(unsafe.Pointer(self.buffer)) }
    if self.error != nil { C.free(unsafe.Pointer(self.error)) }
    C.free(unsafe.Pointer(self))
  }
}

func main() {}