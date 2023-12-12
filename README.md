# enron_mail-indexer

Esta aplicación en go nos permitirá indexar la base de datos de Enron Corp a Zincsearch

# Herramientas necesarias
<ul>
  <li><a href="https://zincsearch-docs.zinc.dev/">Zincsearch</a></li>
  <li>Base de Datos de correos de <a href="http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz">Enron Corp</a></li>
  <li><a href="https://go.dev/">Go</a></li>
  <li><a href="https://graphviz.org/download/https://graphviz.org/download/">Graphviz</a></li>
</ul>

# Ejecutar el programa
<ol>
  <li>Instalar Zincsearch y ejecutar con los siguientes parametros</li>
  <pre>
ZINC_FIRST_ADMIN_USER=admin ZINC_FIRST_ADMIN_PASSWORD=Complexpass#123</pre>
  <li>Descargar la base de datos de correos y guardar la ubicacion del directorio</li>
  <li>Clonar el repositorio (En caso de tener otra version de go hacer lo sgte)</li>
  <pre>
  1. Eliminar los archivos go.mod y go.sum
  2. Abrir el cmd y ejecutar el comando:
      go mod init enron_mail-indexer
      go mod tidy</pre>
  <li>Ejecutar el programa en go y especificar la ruta de la base de datos como parametro:</li>
  <pre>go run main.go full\of\path</pre>
  Esto creará el index Mail y en seguida empezará a indexar los mails en zincsearch
  <li>Para visualizar el grafico del profile hacer lo siguiente: </li>
  <ul>
  <li>Verificar si graphviz esta instalado correctamente:</li>
  <pre>dot -V</pre>
  <li>debera mostrarse lo siguiente o similar:</li>
  <pre>dot - graphviz version 9.0.0 (20230911.1827)</pre>
  <li>Luego, en la ruta del proyecto ejecutar el siguiente comando:</li>
  <pre>go tool pprof -http=:8080 cpu-v1.prof</pre>
  donde cpu-v1.prof es la version del indexador, puede usar cpu-v1.prof o cpu-v2.prof. Este comando le dirigira al navegador con la ruta http://localhost:8080/ui/
  </ul>
</ol>

