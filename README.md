# adhango

adhango is a Go library for calculating Islamic prayer times.

All astronomical calculations are high precision equations directly from the book [“Astronomical Algorithms” by Jean Meeus](https://www.willbell.com/math/mc1.htm). This book is recommended by the Astronomical Applications Department of the U.S. Naval Observatory and the Earth System Research Laboratory of the National Oceanic and Atmospheric Administration.

Implementations of Adhan in other languages can be found in the parent repo [Adhan](https://github.com/batoulapps/Adhan).


## Usage

### Importing
To use this module, import it by running `go get github.com/mnadev/adhango`. View it at [pkg.go.dev](https://pkg.go.dev/github.com/mnadev/adhango).

### Examples

See `example/example.go` to see complete examples of usage.

### Initialization parameters

#### Date

The date parameter passed in should correspond `DateComponents` interface. The year, month, and day values need to be populated. The year, month and day values should be for the local date that you want prayer times for. These date values are expected to be for the Gregorian calendar. The `NewDateComponents` function takes a `time.Time` object and return a `DateComponents` object.

```go
date := data.NewDateComponents(time.Date(2015, time.Month(7), 12, 0, 0, 0, 0, time.UTC))
```

#### Coordinates

Create a `Coordinates` object with the latitude and longitude for the location you want prayer times for.

```go
coords, err := util.NewCoordinates(35.7750, -78.6336) // New York
if err != nil {
    fmt.Printf("got error %+v", err)
    return
}
```

#### Calculation parameters

The rest of the needed information is contained within the `CalculationParameters` interface.

The recommended way to initialize a `CalculationParameters` object is by using the `GetMethodParameters` function and passing in the desired `CalculationMethod`. You can then further customize the calculation parameters if needed. 

```go
params := calc.GetMethodParameters(calc.NORTH_AMERICA)
```

Alternatively, if you want you could use the `CalculationParametersBuilder` to build a `CalculationParameters` object.

```go
params := calc.NewCalculationParametersBuilder().
    SetMadhab(calc.HANAFI).
    SetMethod(calc.NORTH_AMERICA).
    SetFajrAngle(15.0).
    SetIshaAngle(15.0).
    SetMethodAdjustments(calc.PrayerAdjustments{
        DhuhrAdj: 1,
    }).
    Build()
```

##### List of parameters

| Parameter | Description |
| --------- | ----------- |
| `Method`    | CalculationMethod name |
| `FajrAngle` | Angle of the sun used to calculate Fajr |
| `IshaAngle` | Angle of the sun used to calculate Isha |
| `IshaInterval` | Minutes after Maghrib (if set, the time for Isha will be Maghrib plus ishaInterval) |
| `Madhab` | Value from the Madhab enum, used to calculate Asr |
| `HighLatitudeRule` | Value from the HighLatitudeRule enum, used to set a minimum time for Fajr and a max time for Isha |
| `Adjustments` | Struct with custom prayer time adjustments in minutes for each prayer time |
| `MethodAdjustments` | Struct with custom prayer time adjustments in minutes for each prayer time |

**CalculationMethod**

| Value | Description |
| ----- | ----------- |
| `MUSLIM_WORLD_LEAGUE` | Muslim World League. Fajr angle: 18, Isha angle: 17 |
| `EGYPTIAN` | Egyptian General Authority of Survey. Fajr angle: 19.5, Isha angle: 17.5 |
| `KARACHI` | University of Islamic Sciences, Karachi. Fajr angle: 18, Isha angle: 18 |
| `UMM_AL_QURA` | Umm al-Qura University, Makkah. Fajr angle: 18.5, Isha interval: 90. *Note: you should add a +30 minute custom adjustment for Isha during Ramadan.* |
| `DUBAI` | Method used in UAE. Fajr and Isha angles of 18.2 degrees. |
| `MOONSIGHTING_COMMITTEE` | Moonsighting Committee. Fajr angle: 18, Isha angle: 18. Also uses seasonal adjustment values. |
| `NORTH_AMERICA` | Referred to as the ISNA method. This method is included for completeness but is not recommended. Fajr angle: 15, Isha angle: 15 |
| `KUWAIT` | Kuwait. Fajr angle: 18, Isha angle: 17.5 |
| `QATAR` | Modified version of Umm al-Qura used in Qatar. Fajr angle: 18, Isha interval: 90. |
| `SINGAPORE` | Method used by Singapore. Fajr angle: 20, Isha angle: 18. |
| `OTHER` | Fajr angle: 0, Isha angle: 0. This is the default value for `Method` when initializing a `CalculationParameters` object. |

**Madhab**

| Value | Description |
| ----- | ----------- |
| `SHAFI` | Earlier Asr time |
| `HANAFI` | Later Asr time |

**HighLatitudeRule**

| Value | Description |
| ----- | ----------- |
| `MIDDLE_OF_THE_NIGHT` | Fajr will never be earlier than the middle of the night and Isha will never be later than the middle of the night |
| `SEVENTH_OF_THE_NIGHT` | Fajr will never be earlier than the beginning of the last seventh of the night and Isha will never be later than the end of the first seventh of the night |
| `TWILIGHT_ANGLE` | Similar to `SEVENTH_OF_THE_NIGHT`, but instead of 1/7, the fraction of the night used is fajrAngle/60 and ishaAngle/60 |


#### Prayer Times

After that, you can initialize the `PrayerTimes` struct by calling the `NewPrayerTimes` function. The `PrayerTimes` struct will hold timings for `Fajr`, `Sunrise`, `Dhuhr`, `Asr`, `Maghrib` and `Isha` as `time.Time` objects.

By default, the timings are in UTC. In order to specify a specific time zone, call the `SetTimeZone` function on the created `PrayerTimes` struct with the [tz database time zone](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).

```go
prayerTimes, err := calc.NewPrayerTimes(coords, date, params)
if err != nil {
    fmt.Printf("got error %+v", err)
    return
}

err = prayerTimes.SetTimeZone("America/New_York")
if err != nil {
    fmt.Printf("got error %+v", err)
    return
}

fmt.Printf("Fajr: %+v\n", prayerTimes.Fajr)       // Fajr: 2015-07-12 04:42:00 -0400 EDT
fmt.Printf("Sunrise: %+v\n", prayerTimes.Sunrise) // Sunrise: 2015-07-12 06:08:00 -0400 EDT
fmt.Printf("Dhuhr: %+v\n", prayerTimes.Dhuhr)     // Dhuhr: 2015-07-12 13:21:00 -0400 EDT
fmt.Printf("Asr: %+v\n", prayerTimes.Asr)         // Asr: 2015-07-12 18:22:00 -0400 EDT
fmt.Printf("Maghrib: %+v\n", prayerTimes.Maghrib) // Maghrib: 2015-07-12 20:32:00 -0400 EDT
fmt.Printf("Isha: %+v\n", prayerTimes.Isha)       // Isha: 2015-07-12 21:57:00 -0400 EDT
```
