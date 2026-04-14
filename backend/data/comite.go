package data

import "oasis-web/backend/models"

var Comite = map[string]models.Miembro{

	"presidente": {
		Cargo:     "Presidente",
		Nombre:    "Jorge Bouchot",
		Perfil:    "Habilidades de liderazgo, comunicación y capacidad de negociación.",
		Funciones: "Representación oficial del comité y enlace oficial con la comunidad de vecinos. Gestión del grupo de trabajo.",
		Principio: "No toma decisiones solo; siempre actúa bajo el respaldo del acta firmada por la mayoría.",
		Telefono:  "0000000000",
		Lote:      "Lt. 20",
	},

	"secretario": {
		Cargo:     "Secretario",
		Nombre:    "Edgar Hernández",
		Perfil:    "Organizado, con buena redacción y habilidades de conciliación.",
		Funciones: "Redacción de actas de asamblea, control del archivo de vecinos y gestión de reglamento. Apoyo en la gestión del grupo de trabajo y enlace con los vecinos.",
		Principio: "Es quien da validez a los acuerdos. Si no está en el acta, 'no existe', lo que evita malentendidos.",
		Telefono:  "0000000000",
		Lote:      "Lt.",
	},

	"tesoreria": {
		Cargo:     "Tesorera",
		Nombre:    "Elizabeth Calderón",
		Perfil:    "Conocimientos administrativos, habilidades de organización y gestión de cuenta bancaria.",
		Funciones: "Administración de cuenta bancaria, recepción de cuotas de mantenimiento y pago a proveedores. En caso de nuevas necesidades propone nuevas erogaciones.",
		Principio: "Utiliza una cuenta bancaria mancomunada. No toma decisiones de nuevas erogaciones.",
		Telefono:  "0000000000",
		Lote:      "Lt.",
	},

	"analisis": {
		Cargo:     "Vocal de Análisis de Tesorería",
		Nombre:    "Alberto Noble",
		Perfil:    "Capacidad analítica y control financiero.",
		Funciones: "Elaboración de reportes contables mensuales y reporte contable de cierre anual. En caso de nuevas necesidades y proveedores aprueba nuevas erogaciones.",
		Principio: "Reportes públicos oportunos y transparentes para evitar sospechas de mal manejo. No maneja la cuenta bancaria, pero aprueba erogaciones.",
		Telefono:  "0000000000",
		Lote:      "Lt.",
	},

	"seguridad": {
		Cargo:     "Vocal de Seguridad y Accesos I",
		Nombre:    "Erick Pedraza",
		Perfil:    "Enfoque práctico y preventivo con habilidades de informática y telecomunicaciones. Gestión de conflictos con vecinos y proveedores.",
		Funciones: "Supervisión del funcionamiento del portón, mantenimiento de cámaras, gestión de accesos y relación con proveedor de seguridad/portero.",
		Principio: "Función técnica: reportar fallas y proponer mejoras, no confrontar intrusos.",
		Telefono:  "0000000000",
		Lote:      "Lt.",
	},

	"servicios": {
		Cargo:     "Vocal de Gestión de Proveedores",
		Nombre:    "Marisela Bustamante",
		Perfil:    "Disponibilidad para gestión de accesos a proveedores y organización de horarios.",
		Funciones: "Gestión de comunicación con proveedores (agua, basura, CFE, jardinería) y provisión de accesos.",
		Principio: "Gestiona relación con proveedores, no responde por la calidad del servicio.",
		Telefono:  "0000000000",
		Lote:      "Lt.",
	},
}
