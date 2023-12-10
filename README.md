# enron_mail-indexer

Esta aplicación en go nos permitirá indexar la base de datos de Enron Corp a Zincsearch

# Herramientas necesarias
<ul>
  <li><a href="https://zincsearch-docs.zinc.dev/">Zincsearch</a></li>
  <li>Base de Datos de correos de <a href="http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz">Enron Corp</a></li>
  <li><a href="https://go.dev/">Go</a></li>
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
</ol>