1. Search API for box-search, radial-search
2. Implementation of Schema (TimeSpan object)
3. Example queries for normal routes, tests


Did we need:
1. Directions table ??? If not DELETE
2. GTFS Location group should be on UUID i think (no ID)
3. Stations should have UUID
4. Remove GTFS FeedInfo

Agency:
- Add UUID
- Rename ID Agency_id to id PK

Locations -> JSON ->
{
  stations: int, int int,
  regions: int, int, int
}

---- V2
LEARN abiut FREQUENCIES - HOW TO
LEARN about FARES - HOW TO
