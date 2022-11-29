# goaudiofile

## What is it?

It's a music and audio file format support library for Go.

## What formats does it support?

| Subfolder | Name | Notes |
|-----------|------|-------|
| `s3m` | Scream Tracker 3 Module | Based on the format described in `TECH.DOC`, originally supplied with the Scream Tracker 3 application, by Sami Tammilehto / FutureCrew |
| `mod` | Protracker / Fast Tracker Module | Based on the format described in `FMODDOC.TXT`, originally supplied with the FireMOD 1.06 source code distribution, by Brett Paterson / FireLight. In order to stay free of copyright concerns (FireLight still operates and maintains FMOD / FireMOD), the associated FireMOD source code was not referenced during the creation of this library. Any similarities of this library to the FireMOD source code is purely accidental and coincidental. |

## Bugs

### Known Bugs

| Tags | Notes |
|------|-------|
| `s3m` | The technical document describing the S3M format has many errors and inconsistencies that have been speculated and argued over by many experts in the field for many decades. This implementation attempts to use the least troublesome representation of each point, where possible. As a result, the data obtained from a format read with this library might not produce a 100% accurate-to-ST3 result. |
| `mod` | If you thought `s3m` was a truly-inconsistent format, then you obviously haven't met its older brother, the Protracker/FastTracker `mod`. |
