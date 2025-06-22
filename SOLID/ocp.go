/*
Solid principles - Open/Closed Principle (OCP)
The Open/Closed Principle states that software entities (classes, modules, functions, etc.
should be open for extension but closed for modification. 
This means that you should be able to add new functionality to existing code
without changing the existing codebase.
This principle helps in maintaining the stability of the code 
while allowing for new features to be added without breaking existing functionality.
In Go, we can apply OCP by using interfaces and embedding structs
to extend functionality without modifying the existing code.
*/

package solid

func ExportSurvey(s * Survey, exporter Exporter) error {
	return exporter.Export(s)
}

type Exporter interface {
	Export(s *Survey) error
}

type S3Exporter struct {}

func (e *S3Exporter) Export(s *Survey) error {
 return nil
}

type GCSExporter struct {}

func (e *GCSExporter) Export(s *Survey) error {
 return nil
}
