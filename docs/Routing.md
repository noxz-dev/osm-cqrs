# Routing

Recherche ergab, dass bei Änderungen in der OSM-Datenbank die Routing-Datenbank neu berechnet wird. Gefundene kompatible Routing-Bibliotheken ermöglichen keine inkrementellen Aktualisierung.

Daher muss vermutlich untersucht werden, wie lange es dauert diese neu zu generieren.

Vielleicht lassen sich auch Bounding Boxen definieren, sodass nur dieser Bereich innerhalb der Routing-Datenbank aktualisiert werden muss.

Vermutlich muss die Routing Datenbank nur aktualisiert werden, wenn auch explizit Kanten welche für das Routing erstellt, verändert und gelöscht werden. Muss analysiert werden, welche Tags relevant dafür sind.

## Bibliotheken

### [pgRouting](https://pgrouting.org/)

Erweiterung for PostgreSQL / PostGIS.

- Sehr flexibel, weil SQL-basiert
- Langsam

#### [osm2pgrouting](https://github.com/pgRouting/osm2pgrouting)

Ermöglicht es `.osm`-Daten automatisch in das von pgRouting benötigte Format zu konvertieren und entsprechend in die Datenbank einfügt.

**Noch keine Lösung für `OsmChange`-Daten gefunden**. Mittels `osm2pgrouting` lassen sich `osm`-Daten inkrementell hinzufügen. Bietet aber keine expliziete Lösung für das Updaten von Daten.

### [Graphhopper](https://github.com/graphhopper/graphhopper)

- Sehr schnell
- Kann als Bibliothek und API verwendet werden

### [OSRM](https://github.com/Project-OSRM/osrm-backend)

- Standard-Routing von OSM
- schnell
- HTTP API, C++ Lib, NodeJs wrapper
