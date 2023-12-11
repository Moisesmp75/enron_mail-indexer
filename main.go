package main

import (
	"enron_mail-indexer/modules"
	"fmt"
	"os"
)

// var (
// 	cpuprofile = flag.String("cpuprofile", "", "escribe el perfil de la CPU en `file`")
// 	memprofile = flag.String("memprofile", "", "escribe el perfil de memoria en `file`")
// )

func main() {
	// flag.Parse()
	// if *cpuprofile != "" {
	// 	f, err := os.Create(*cpuprofile)
	// 	if err != nil {
	// 			log.Fatal("no se pudo crear el perfil de CPU: ", err)
	// 	}
	// 	defer f.Close()
	// 	if err := pprof.StartCPUProfile(f); err != nil {
	// 			log.Fatal("no se pudo iniciar el perfil de CPU: ", err)
	// 	}
	// 	defer pprof.StopCPUProfile()
	// }
	// modules.Create_index_zincsearch()

	if len(os.Args) < 2 {
		fmt.Println("Falta el path del directorio como argumento")
		return
	}
	path := os.Args[1]
	fmt.Println(path)
	modules.IndexerV2(path)
	// modules.IndexerConcurrent("D:\\Descargas\\enron_mail_20110402\\enron_mail_20110402\\maildir",runtime.NumCPU())

	// if *memprofile != "" {
	// 	f, err := os.Create(*memprofile)
	// 	if err != nil {
	// 			log.Fatal("no se pudo crear el perfil de memoria: ", err)
	// 	}
	// 	defer f.Close()
	// 	runtime.GC()
	// 	if err := pprof.WriteHeapProfile(f); err != nil {
	// 			log.Fatal("no se pudo escribir el perfil de memoria: ", err)
	// 	}
	// }
	// fmt.Println("Tiempo transcurrido:", duration)

}
