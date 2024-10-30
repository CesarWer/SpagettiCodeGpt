package Services

/*
import (
	"Spaghetti-Code/Models"
	"Spaghetti-Code/connections"
	"Spaghetti-Code/dto"


	"Spaghetti-Code/Models"
	"Spaghetti-Code/dto"
	//"Spaghetti-Code/repositories"

	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

)

var compraCollection *mongo.Collection = connections.GetCollection("compras")




//var compraCollection *mongo.Collection = repositories.GetCollection("compras")
/*

func ListarProductosBajoStock(tipo string, nombre string) ([]Models.Alimento, error) {
	var alimentos []Models.Alimento

	filtro := bson.M{
		"cantidad_actual": bson.M{"$lt": bson.M{"$expr": "$cantidad_minima"}},
	}

	if tipo != "" {
		filtro["tipo"] = bson.M{"$regex": tipo, "$options": "i"}
	}

	if nombre != "" {
		filtro["nombre"] = bson.M{"$regex": nombre, "$options": "i"}
	}

	cursor, err := alimentoCollection.Find(context.TODO(), filtro)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &alimentos); err != nil {
		return nil, err
	}

	return alimentos, nil
}

func CrearCompra() (dto.CompraDTO, error) {
	var alimentos []Models.Alimento
	var compraDTO dto.CompraDTO
	var totalCosto float64

	// Buscar alimentos con cantidad menor a la mínima
	filtro := bson.M{
		"cantidad_actual": bson.M{"$lt": bson.M{"$expr": "$cantidad_minima"}},
	}

	cursor, err := alimentoCollection.Find(context.TODO(), filtro)
	if err != nil {
		return compraDTO, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &alimentos); err != nil {
		return compraDTO, err
	}

	// Calcular el costo total y actualizar las cantidades
	for _, alimento := range alimentos {
		cantidadRequerida := alimento.CantidadMinima - alimento.CantidadActual
		costoTotalProducto := float64(cantidadRequerida) * alimento.PrecioUnitario

		// Crear el producto para el DTO
		productoCompra := dto.ProductoCompraDTO{
			AlimentoID:       alimento.ID.Hex(),
			Nombre:           alimento.Nombre,
			CantidadComprada: cantidadRequerida,
			PrecioUnitario:   alimento.PrecioUnitario,
			CostoTotal:       costoTotalProducto,
		}

		// Sumar al costo total
		totalCosto += costoTotalProducto

		// Actualizar la cantidad en la base de datos
		_, err := alimentoCollection.UpdateOne(
			context.TODO(),
			bson.M{"_id": alimento.ID},
			bson.M{
				"$set": bson.M{"cantidad_actual": alimento.CantidadMinima},
			},
		)
		if err != nil {
			return compraDTO, err
		}

		compraDTO.Productos = append(compraDTO.Productos, productoCompra)
	}

	// Crear la compra final
	compraDTO.Fecha = time.Now()
	compraDTO.CostoTotal = totalCosto

	// Insertar la compra en la base de datos
	_, err = compraCollection.InsertOne(context.TODO(), compraDTO)
	if err != nil {
		return compraDTO, err
	}

	return compraDTO, nil
}*/
