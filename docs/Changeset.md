# Changeset

## Dokumetation

[OSM-Wiki - Changeset](https://wiki.openstreetmap.org/wiki/Changeset)

Dient an sich nur als Gruppierung von einer Anzahl an Änderungen. Enthalten keine konkreten Änderungen. Dieser werden in _OsmChange_ Objekten dokumentiert. Enthält Informationen darüber wer die Änderungen vorgenommmen in welchem Bereich und mit welchem Programm.

## Beispiele

[Beispiel eines Änderungssatzes visualisiert in OSM](https://www.openstreetmap.org/changeset/97929008)

[XML-Änderungssatz](https://www.openstreetmap.org/api/0.6/changeset/97929008)

# OsmChange

## Dokumentation

[OSM-Wiki - OsmChange](https://wiki.openstreetmap.org/wiki/OsmChange)

## Beispiele

[Konkreten Änderungen im Änderungssatz](https://www.openstreetmap.org/api/0.6/changeset/97929008/download)

Es scheint so als wenn in den `create` und `modify` Elementen immer das gesamte Element mitgesendet wird. Vor allem bei `modify` wird entsprechend immer die aktuellste Version mitgeschickt. Wurde beispielsweise eine Relation bearbeitet werde alle Nodes die bereits vorhanden waren plus den neuen Node mitgesendet.
Bei Elementen innerhalb eines `delete`-Elements ist der Tag `visible` auf `false`. Wenn dieser auf `false` ist darf dieser in keiner weiteren Operation betrachtet werden, da dieser in OSM auch nicht mehr angezeigt wird.

**Wird immer die neueste Version in `modify` gesendet oder gibt es inkrementelle updates?**

## Vorverarbeitung von Daten

### [osmconvert](https://wiki.openstreetmap.org/wiki/Osmconvert)

Tool was das Konvertieren von Daten ermöglicht. Explizit auch entsprechende Tags herausfiltern.
