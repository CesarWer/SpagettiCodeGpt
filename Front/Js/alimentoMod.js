document.addEventListener("DOMContentLoaded", function() {
    // Obtener el ID del alimento de la URL
    const urlParams = new URLSearchParams(window.location.search);
    const alimentoId = urlParams.get('id');

    if (alimentoId) {
        // Realizar solicitud para obtener los datos del alimento
        makeRequest(
            `http://localhost:8080/api/alimentos/${alimentoId}`,
            Method.GET,
            null,
            ContentType.JSON,
            CallType.PRIVATE,
            (response) => {
                // Llenar los campos del formulario con la respuesta
                document.getElementById("nombre").value = response.nombre;
                document.getElementById("tipo").value = response.tipo;
                document.getElementById("momentos").value = response.momentos;
                document.getElementById("cantidadActual").value = response.cantidadActual;
                document.getElementById("cantidadMinima").value = response.cantidadMinima;
                document.getElementById("precioUnitario").value = response.precioUnitario;
            },
            (status, error) => {
                console.error("Error al obtener el alimento:", status, error);
            }
        );
    }
});


function modificarAlimento() {
    // Obtener el ID del alimento de la URL
    const urlParams = new URLSearchParams(window.location.search);
    const alimentoId = urlParams.get('id');

    // Obtener los valores del formulario
    const data = {
        nombre: document.getElementById("nombre").value,
        tipo: document.getElementById("tipo").value,
        momentos: document.getElementById("momentos").value,
        cantidadActual: parseInt(document.getElementById("cantidadActual").value),
        cantidadMinima: parseInt(document.getElementById("cantidadMinima").value),
        precioUnitario: parseFloat(document.getElementById("precioUnitario").value)
    };

    // Realizar solicitud para modificar el alimento
    makeRequest(
        `http://localhost:8080/api/alimentos/${alimentoId}`,
        Method.PUT,
        data,
        ContentType.JSON,
        CallType.PRIVATE,
        (response) => {
            alert("Alimento modificado exitosamente!");
            // Redirigir a la página de alimentos
            window.location.href = "alimentos.html";
        },
        (status, error) => {
            console.error("Error al modificar el alimento:", status, error);
        }
    );
}
