package htmlutil

import (
	"github.com/microcosm-cc/bluemonday"
	"html"
	"reflect"
)


func SanitizeJSON(v interface{}) interface{} {

	newV := translate(v)

	return newV
}


func SanitizeString(str string) string {

	newStr := translate(str)

	return newStr.(string)
}


func translate(obj interface{}) interface{} {
	// Wrap the original in a reflect.Value
	original := reflect.ValueOf(obj)

	copy := reflect.New(original.Type()).Elem()
	translateRecursive(copy, original)

	// Remove the reflection wrapper
	return copy.Interface()
}

func translateRecursive(copy, original reflect.Value) {
	switch original.Kind() {
	// The first cases handle nested structures and translate them recursively

	// If it is a pointer we need to unwrap and call once again
	case reflect.Ptr:
		// To get the actual value of the original we have to call Elem()
		// At the same time this unwraps the pointer so we don't end up in
		// an infinite recursion
		originalValue := original.Elem()
		// Check if the pointer is nil
		if !originalValue.IsValid() {
			return
		}
		// Allocate a new object and set the pointer to it
		copy.Set(reflect.New(originalValue.Type()))
		// Unwrap the newly created pointer
		translateRecursive(copy.Elem(), originalValue)

	// If it is an interface (which is very similar to a pointer), do basically the
	// same as for the pointer. Though a pointer is not the same as an interface so
	// note that we have to call Elem() after creating a new object because otherwise
	// we would end up with an actual pointer
	case reflect.Interface:
		// Get rid of the wrapping interface
		originalValue := original.Elem()

		if originalValue.IsValid() {
			// Create a new object. Now new gives us a pointer, but we want the value it
			// points to, so we have to call Elem() to unwrap it
			copyValue := reflect.New(originalValue.Type()).Elem()
			translateRecursive(copyValue, originalValue)
			copy.Set(copyValue)
		}


	// If it is a struct we translate each field
	case reflect.Struct:
		for i := 0; i < original.NumField(); i += 1 {
			translateRecursive(copy.Field(i), original.Field(i))
		}

	// If it is a slice we create a new slice and translate each element
	case reflect.Slice:
		copy.Set(reflect.MakeSlice(original.Type(), original.Len(), original.Cap()))
		for i := 0; i < original.Len(); i += 1 {
			translateRecursive(copy.Index(i), original.Index(i))
		}

	// If it is a map we create a new map and translate each value
	case reflect.Map:
		copy.Set(reflect.MakeMap(original.Type()))
		for _, key := range original.MapKeys() {
			originalValue := original.MapIndex(key)
			if originalValue.IsNil() {
				continue
			}
			// New gives us a pointer, but again we want the value
			copyValue := reflect.New(originalValue.Type()).Elem()

			translateRecursive(copyValue, originalValue)
			copy.SetMapIndex(key, copyValue)
		}

	// Otherwise we cannot traverse anywhere so this finishes the the recursion

	// If it is a string translate it (yay finally we're doing what we came for)
	case reflect.String:

		p := bluemonday.UGCPolicy()

		elements := []string{"img", "i", "a", "tr", "td", "table", "tbody", "thead", "span", "div", "p", "li", "ul", "ol", "b", "strong", "hr", "h1", "h2", "h3", "h4", "h5", "label", "small", "font"}

		p.AllowAttrs("style", "class").OnElements(elements...)
		p.AllowElements(elements...)

		newStr := html.UnescapeString(p.Sanitize(original.Interface().(string)))

		copy.SetString(newStr)

	// And everything else will simply be taken from the original
	default:
		copy.Set(original)
	}

}
