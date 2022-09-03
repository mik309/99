# Prueba Tecnica 99minutos.com
Prueba desarrollada en golang y gingonic.

Base de datos utilizada: Postgresql

## Instalación 

_Solo es necesario hacer un gitclone de la main branch_

```
git clone https://github.com/mik309/99.git
```

### Iniciar proyecto
_Solo debemos ejecutar el comando_

```
go run .
```

#### Funcionalidad

Esta API cuenta con 4 endpoints

**Crear Usuario** - Nos permite crear un usuario [TODOS TIENEN ACCESO]. Es necesario:
  
  * FirstName - string
  * LastName - string
  * IsAdmin - bool
  * Email - string
  * Password - string
  
**Crear Orden** - Nos permite crear una orden [CLIENTES TIENEN ACCESO]. Es necesario:

  [CONTIENE UN BASIC AUTH]
  * DestinationAddres
      * FirstName - string
      * LastName - string
      * Street - string
      * ZipCode - string
      * State - string
      * City - string
      * Neighbourhood - string
      * ExNumber - string
      * PhoneNumber - string
  * Products - List of products
      * Weight - float

** Obtener Orden - Permite la busqueda de una order [CLIENTES TIENEN ACCESO]. Es necesario:

[CONTIENE UN BASIC AUTH]
  * Debemos pasar en el path el id de la orden - /id
  
** Actualizar Estatus - Permite actualizar el estatus de una orden [ADMINISTRADORES TIENEN ACCESO]. Es necesario:

[CONTIENE UN BASIC AUTH]
  * Debemos pasar en el path la id seguida del nuevo estatus - /id/nuevo_estatus
  
      
      
