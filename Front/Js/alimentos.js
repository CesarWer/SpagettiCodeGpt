function cargarAlimentos() {
    const usuarioId = 1;
    const url = `http://localhost:8080/alimentos/${usuarioId}`;
    makeRequest(
      url,              // URL del endpoint
      Method.GET,                     // Método HTTP
      null,                           // Sin datos para GET
      ContentType.JSON,               // Tipo de contenido
      CallType.PUBLIC,                // Llamada pública o privada según tu configuración
      (response) => {                 // Callback de éxito
        const tableBody = document.getElementById('alimentos-table-body');
        tableBody.innerHTML = ''; // Limpia cualquier contenido previo
  
        response.forEach(alimento => {
          const row = document.createElement('tr');
  
          // Crear la fila con los datos del alimento
          row.innerHTML = `
            <td>${alimento.nombre}</td>
            <td>${alimento.tipo}</td>
            <td>${alimento.momentos.join(', ')}</td>
            <td>${alimento.cantidadActual}</td>
            <td>${alimento.cantidadMinima}</td>
            <td>$${alimento.precioUnitario.toFixed(2)}</td>
            <td>
              <button class="btn btn-warning btn-sm" onclick="modificarAlimento('${alimento.id}')">Modificar</button>
              <button class="btn btn-danger btn-sm" onclick="eliminarAlimento('${alimento.id}')">Eliminar</button>
            </td>
          `;
  
          tableBody.appendChild(row);
        });
      },
      (status, error) => {            // Callback de error
        console.error(`Error ${status}: ${error}`);
      }
    );
  }
  
  // Llamar a cargarAlimentos al cargar la página
  document.addEventListener('DOMContentLoaded', cargarAlimentos);
  
  // Función para modificar alimento (puedes hacer una solicitud PUT)
  function modificarAlimento(alimentoId) {
    // Redirige a modificar.html pasando el ID del alimento como parámetro
    window.location.href = `alimentoMod.html?id=${alimentoId}`;
}

  
  // Función para eliminar alimento (podrías hacer una solicitud DELETE)
  function eliminarAlimento(id) {
    if (confirm("¿Estás seguro de que deseas eliminar este alimento?")) {
      console.log(`Eliminando alimento con ID: ${id}`);
      makeRequest(
        `/api/alimentos/${id}`,  // Endpoint de eliminación
        Method.DELETE,           // Método DELETE
        null,                    // No necesitas datos en el cuerpo
        ContentType.JSON,        // Tipo de contenido
        CallType.PRIVATE,        // Asegúrate de que la llamada sea privada si es necesario
        () => {                  // Callback de éxito
          alert("Alimento eliminado correctamente");
          cargarAlimentos();     // Recargar los alimentos después de la eliminación
        },
        (status, error) => {     // Callback de error
          alert(`Error al eliminar alimento: ${error.message || status}`);
        }
      );
    }
  }

  
  // Función para crear un nuevo alimento
  function crearAlimento() {
    // Obtener los valores de los campos
    const nombre = document.getElementById('nombre').value;
    const tipo = document.getElementById('tipo').value;
    const momentos = document.getElementById('momentos').value.split(',').map(m => m.trim());
    const cantidadActual = parseInt(document.getElementById('cantidadActual').value);
    const cantidadMinima = parseInt(document.getElementById('cantidadMinima').value);
    const precioUnitario = parseFloat(document.getElementById('precioUnitario').value);
    const usuarioID = 1;

    if (isNaN(cantidadActual) || isNaN(cantidadMinima) || isNaN(precioUnitario)) {
      alert("Por favor, ingresa valores numéricos válidos para cantidad y precio.");
      return;
    }  
    // Crear el objeto con los datos del alimento
    const nuevoAlimento = {
      nombre,
      tipo,
      momentos,
      cantidadActual,
      cantidadMinima,
      precioUnitario,
      usuarioID
    };
  
    // Enviar la solicitud al backend
    makeRequest(
      "http://localhost:8080/alimentos/", // URL completa al endpoint de creación
      Method.POST,
      nuevoAlimento,
      ContentType.JSON,
      CallType.PRIVATE,
      () => {
        alert("Alimento creado con éxito");
        cargarAlimentos(); // Recargar la lista de alimentos
        document.getElementById('crear-alimento-form').reset();
      },
      (status, error) => {
        alert(`Error al crear alimento: ${error.message || status}`);
      }
    );
  }
  
  
  