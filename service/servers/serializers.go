package servers

import (
	"github.com/peake100/stalkforecaster-go/protogen/stalk_proto"
	"github.com/peake100/turnup-go/models"
)

func serializePricePeriods(pricePeriods []models.PricePeriod) []int32 {
	serialized := make([]int32, len(pricePeriods))
	for i, period := range pricePeriods {
		serialized[i] = int32(period)
	}
	return serialized
}

func serializePriceSeries(series *models.PriceSeries) *stalkproto.PricesSummary {
	return &stalkproto.PricesSummary{
		Min:               int32(series.MinPrice()),
		Guaranteed:        int32(series.GuaranteedPrice()),
		Max:               int32(series.MaxPrice()),
		MinPeriods:        serializePricePeriods(series.MinPeriods()),
		MaxPeriods:        serializePricePeriods(series.MaxPeriods()),
		GuaranteedPeriods: serializePricePeriods(series.GuaranteedPeriods()),
	}
}

func serializeBreakdown(breakdown *models.SpikeChanceBreakdown) []float64 {
	breakdownSerialized := make([]float64, len(breakdown))
	for i, chance := range breakdown {
		breakdownSerialized[i] = chance
	}
	return breakdownSerialized
}

func serializeForecastSpikes(spike models.HasSpikeChance) *stalkproto.SpikeChances {
	return &stalkproto.SpikeChances{
		Has:       spike.Has(),
		Start:     int32(spike.Start()),
		End:       int32(spike.Start()),
		Chance:    spike.Chance(),
		Breakdown: serializeBreakdown(spike.Breakdown()),
	}
}

func serializeSpikeRange(spike models.HasSpikeRange) *stalkproto.SpikeRange {
	return &stalkproto.SpikeRange{
		Has:   spike.Has(),
		Start: int32(spike.Start()),
		End:   int32(spike.End()),
	}
}

func serializeWeek(
	week *models.PotentialWeek, pattern models.PricePattern,
) *stalkproto.PotentialWeek {
	prices := make([]*stalkproto.PricePeriod, 12)
	for i, thisPrice := range week.Prices {
		prices[i] = &stalkproto.PricePeriod{
			Min:     int32(thisPrice.MinPrice()),
			Max:     int32(thisPrice.MaxPrice()),
			IsSpike: thisPrice.Spikes.Any().Has(),
		}
	}

	return &stalkproto.PotentialWeek{
		Chance:        week.Chance(),
		Prices:        prices,
		PricesSummary: serializePriceSeries(&week.PriceSeries),
		PricesFuture:  serializePriceSeries(&week.Future),
		// We can use any spike here because in this instance is will be identical to
		// to whatever spike pattern type we are in
		Spike: serializeSpikeRange(week.Spikes.Any()),
	}
}

func serializePattern(in *models.PotentialPattern) *stalkproto.PotentialPattern {

	weeks := make([]*stalkproto.PotentialWeek, 0)
	for _, thisWeek := range in.PotentialWeeks {
		weekSerialized := serializeWeek(thisWeek, in.Pattern)
		weeks = append(weeks, weekSerialized)
	}

	return &stalkproto.PotentialPattern{
		Pattern:       stalkproto.PricePatterns(in.Pattern),
		Chance:        in.Chance(),
		PricesSummary: serializePriceSeries(&in.PriceSeries),
		PricesFuture:  serializePriceSeries(&in.Future),
		// We can use any spike here because in this instance is will be identical to
		// to whatever spike pattern type we are in
		Spike:          serializeSpikeRange(in.Spikes.Any()),
		PotentialWeeks: weeks,
	}
}

func serializePrediction(prediction *models.Prediction) *stalkproto.Forecast {
	// There are always 4 patterns
	pricePatterns := make([]*stalkproto.PotentialPattern, 4)
	for i, pattern := range prediction.Patterns {
		pricePatterns[i] = serializePattern(pattern)
	}

	forecast := &stalkproto.Forecast{
		PricesSummary: serializePriceSeries(&prediction.PriceSeries),
		PricesFuture:  serializePriceSeries(&prediction.Future),
		Patterns:      pricePatterns,
		Spikes: &stalkproto.ForecastSpikes{
			Small: serializeForecastSpikes(prediction.Spikes.Small()),
			Big:   serializeForecastSpikes(prediction.Spikes.Big()),
			Any:   serializeForecastSpikes(prediction.Spikes.Any()),
		},
		Heat: int32(prediction.Heat),
	}

	return forecast
}
