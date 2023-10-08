package models

var PlanetMap = map[string]Planet{
	"Mercury": MercuryPlanet,
	"Venus":   VenusPlanet,
	"Earth":   EarthPlanet,
	"Mars":    MarsPlanet,
	"Jupiter": JupiterPlanet,
	"Saturn":  SaturnPlanet,
	"Uranus":  UranusPlanet,
	"Neptune": NeptunePlanet,
	"Moon":    MoonPlanet,
}

var MercuryPlanet = Planet{
	Type:                   "Rocky planet",
	Size:                   4880,
	RelatedSizeOfEarth:     38,
	Mass:                   "3.301 × 10^23",
	RealtedMassOfEarth:     0.05,
	LengthOfAYear:          88,
	NumberOfMoons:          0,
	AverageDistanceFromSun: 57.9, // million
	Temperature:            "-173~427",
	Formation: `Like all the other planets, Mercury was likely formed in the large cloud of gas, dust and ice of the ancient solar system, which collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc.
	Mercury is the closest planet to the Sun, which likely influenced its composition and formation. The planet has a surprisingly large iron core, which may have formed because the early Sun heated and vaporized rock on Mercury's surface.`,
	OribitAndRotation: `Named after the Roman messenger of the gods, who was known for being very swift, Mercury takes only 88 days to complete one orbit around the Sun – the shortest year of all planets.

	Just like Earth, Mercury also rotates on its own axis, but much more slowly: one day-night cycle on Mercury takes the equivalent of 176 Earth days, or two full Mercurian years! Mercury’s rotation in relation to distant stars takes 59 Earth days.

	Mercury's orbit is also the least circular of all the planets and looks somewhat egg-shaped. Because of this, the morning Sun appears to rise briefly, set, and continue rising again from certain spots on the surface of the small planet. The opposite effect happens at sunset, with the Sun setting and briefly rising before finally setting completely.`,
	Surface: `At first glance Mercury looks a lot like Earth's Moon. Its surface is littered with craters that have accumulated and remained unchanged over billions of years, since the planet has no significant geological activity or atmosphere.

	Mercury's surface remains the most mysterious of all the rocky planets, though, as no lander has ever explored it. Some of Mercury's craters have shaded sections that are never exposed to the Sun. Scientists believe that frozen water may exist in these craters.`,
	Exosphere: `Mercury's atmosphere is so thin and fragile that it is instead considered an exosphere: a very thin volume of matter that surrounds a planet, but is not dense enough to behave like an atmosphere. For the most part, this outer layer is made up of particles coming from solar winds and elements from Mercury's crust (such as hydrogen, helium, oxygen and sodium) that are vaporized by the intense heat of the Sun.

	Mercury's exosphere is constantly evolving: while the solar wind continuously blows it away, intense gusts from the Sun also constantly super-heat more of the planet's crust, thereby replenishing Mercury's gassy envelope.`,
}

var VenusPlanet = Planet{
	Type:                   "Rocky planet",
	Size:                   12104,
	RelatedSizeOfEarth:     95,
	Mass:                   "4.868 × 10^24",
	RealtedMassOfEarth:     82,
	LengthOfAYear:          224.7,
	NumberOfMoons:          0,
	AverageDistanceFromSun: 108208000,
	Temperature:            "460",
	Formation: `Like all the other planets, Venus was likely formed in the large cloud of gas, dust, and ice of the ancient solar system, which collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc.

	Venus is the second closest planet to the Sun and is often called Earth's "sister planet" or "twin planet." Venus and Earth have very similar sizes, masses, and compositions, but are otherwise quite different from each other. Scientists think that in the early days of the solar system, Venus may have been similar to present-day Earth and may even have had liquid water on its surface.`,
	OribitAndRotation: `Venus is the only planet that rotates clockwise; all the other planets rotate counter-clockwise. Venus's rotation is called retrograde rotation.

	One day on Venus – the time it takes the planet to complete a full rotation on its axis – is equal to 243 days on Earth. That is actually slower than the time it takes for Venus to complete one orbit around the Sun, which takes 225 Earth days.`,
	Surface: `Venus shows evidence of very strong volcanic activity and likely has more volcanoes than our planet. Scientists have never seen an active lava flow on Venus, but they estimate that 80% of the planet's surface is covered by smooth, volcanic plains that are left over from previous lava flows.

	Venus is entirely cloaked in thick clouds that hide its surface. In order to study it, scientists must use either spacecraft that land on the planet or radar technology on board satellites or probes.`,
	Exosphere: `Despite being only the second planet from the Sun, Venus (not Mercury) is the hottest planet in our solar system.

	Venus's high temperature is caused by its extremely thick atmosphere, which is over 90 times thicker than Earth's! This outer layer is composed almost entirely of carbon dioxide, a well-known gas responsible for a very strong greenhouse effect. Trapped by Venus's atmosphere, the Sun's energy heats the planet to constant temperatures over 450 degrees Celsius.`,
}

var EarthPlanet = Planet{
	Type:                   "Rocky planet",
	Size:                   12742,
	RelatedSizeOfEarth:     0,
	Mass:                   "5.97 × 10^24",
	RealtedMassOfEarth:     0.0,
	LengthOfAYear:          365.26,
	NumberOfMoons:          1,
	AverageDistanceFromSun: 150,
	Temperature:            "-89~56.7",
	Formation:              `Scientists believe it took between 10 and 20 million years for Earth to form. The solar system began as a large cloud of gas, dust, and ice that collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc.`,
	OribitAndRotation: `Earth completes one orbit around the Sun in 365.26 days. Since this is a little longer than a regular calendar year of 365 days, one extra day is added every four years; a year with 366 days is called a leap year.
	Earth also spins on its axis and completes one rotation per 24 hours. This rotation causes night and day.
	`,
	Seasons: `Earth's rotation and orbit are not perfectly perpendicular to each other – an angle of 23.5 degrees exists between the two. This tilt in Earth's rotation causes different parts of the planet's surface to receive varying intensities of sunlight throughout the year, causing the seasons.

	When the northern hemisphere is tilted towards the Sun, it is summer in the North and winter in the South. When the southern hemisphere is tilted towards the Sun, it is summer in the South and winter in the North.`,
	Surface:   "Over two-thirds of Earth's surface is covered in water. The abundance of this important substance plays a crucial role in making Earth the only place in the entire universe that we know with absolute certainty harbours life.",
	Exosphere: "",
}

var MarsPlanet = Planet{
	Type:                   "Rocky planet",
	Size:                   6791,
	RelatedSizeOfEarth:     0.5,
	Mass:                   "6.417 × 10^23",
	RealtedMassOfEarth:     0.1,
	LengthOfAYear:          687,
	NumberOfMoons:          2,
	AverageDistanceFromSun: 227939200,
	Temperature:            "-140~30",
	Formation:              `Mars is the fourth planet from the Sun in our solar system. Scientists believe that all of the planets were created just over 4.5 billion years ago. The solar system began as a large cloud of gas, dust, and ice, which collapsed into a spinning disc. The Sun was formed at its centre and particles began sticking together along rings in the disc – leading to the formation of the planets.`,
	OribitAndRotation: `A year on Mars – the time it takes for the planet to orbit the Sun – is nearly twice as long as a year on Earth. However, the planets rotate at a similar frequency: a day on Mars (known as a sol) lasts about 24 hours and 40 minutes in Earth time.

	Mars is tilted by about 25 degrees, which means that the red planet has four seasons, too – each twice as long as Earth's!`,
	Seasons: `Mars is known as the "red planet" due to its reddish hue caused by oxidized iron (or rust) on the planet's surface.

	The tallest mountain in the entire solar system, Olympus Mons, is on Mars. This extinct volcano is nearly 22 km tall – roughly two and a half times the height of Mount Everest.

	Valles Marineris is a massive canyon system that stretches over 3,000 km across Mars's surface – about the distance between Montreal and Calgary! Plunging about 8 km deep, and often referred to as "the Grand Canyon of Mars," Valles Marineris makes the red planet home to the largest known canyon in the solar system.

	`,
	Surface:   "",
	Exosphere: "",
}

var JupiterPlanet = Planet{
	Type:                   "Gas giant",
	Size:                   139822,
	RelatedSizeOfEarth:     11,
	Mass:                   "1.9 × 10^27",
	RealtedMassOfEarth:     318,
	LengthOfAYear:          11.8,
	NumberOfMoons:          79,
	AverageDistanceFromSun: 778.6,
	Temperature:            "-161~-108",
	Formation: `Jupiter is the fifth planet from the Sun and also the largest in our solar system. Like all the other planets, it was likely formed in a massive, ancient cloud of gas, dust, and ice that collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc.

	Scientists believe that many rocky planets larger than Earth existed in the early inner solar system. Physical forces moved Jupiter closer to the Sun and destroyed these other planets, allowing the formation of the current rocky planets (Mercury, Venus, Earth and Mars) before pushing Jupiter back outwards to its current position.

	Jupiter has at least 79 known moons, which were drawn by the massive planet's strong gravitational pull. Fifty-three of them have names, while 26 are not yet named.`,
	OribitAndRotation: `Jupiter has the shortest day of any planet in the solar system: it completes a full rotation on itself in less than 10 hours. Because of its very fast rotation, Jupiter is noticeably flattened at its poles and bulges around its centre.

	Although Jupiter's days are short, its years are long – the planet completes an orbit of the Sun in just under 12 Earth years.`,
	Seasons: "",
	Surface: `Since Jupiter and Saturn are mainly composed of gases instead of solid elements, it can be tricky to determine where their atmosphere ends and their surface begins. However, the surface is defined as the layer with atmospheric pressure equal to that at sea level on Earth.

	Jupiter's atmosphere is very thick and is mostly composed of molecular hydrogen, helium and methane. Cloud layers and swirling storms can be seen clearly in its atmosphere, even with amateur telescopes. One storm, known as the Great Red Spot, has been observed by astronomers for centuries. It is several times bigger than Earth and appears to be shrinking.

	Scientists believe Jupiter may have a small rocky core that could be about the size of Earth, but this has not yet been proven.`,
	Exosphere: "",
}

var SaturnPlanet = Planet{
	Type:                   "Gas giant",
	Size:                   116464,
	RelatedSizeOfEarth:     9,
	Mass:                   "5.68 × 10^26",
	RealtedMassOfEarth:     95,
	LengthOfAYear:          29.5,
	NumberOfMoons:          146,
	AverageDistanceFromSun: 1.43,
	Temperature:            "-189~-139",
	Formation: `Saturn is the sixth planet from the Sun and the second largest in our solar system. Like all the other planets, Saturn was likely formed in the large cloud of gas, dust, and ice which collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc.

	Saturn is best known for its impressive rings, which are about 20 metres thick on average. Some scientists believe that the planet's icy rings are made of material left over from a massive collision between the young planet and one of its early moons.`,
	OribitAndRotation: `Saturn has a very short day. The planet takes about 10.5 hours to complete a full rotation on its axis. It takes Saturn about 29.5 Earth years to complete a full orbit around the Sun.`,
	Seasons:           "",
	Surface: `Since gas giant planets like Saturn and Jupiter are mainly composed of gases instead of solid elements, it can be tricky to determine where their atmosphere ends and their surface begins. However, the surface is defined as the layer with atmospheric pressure equal to that at sea level on Earth.

	There is some debate about the conditions within the planet's atmosphere. For example, some scientists argue that lightning storms in the upper layers even cause it to rain diamonds on Saturn!`,
	Exosphere: "",
	Else: `Moons and rings - Saturn has an abundance of moons. For example, much like Jupiter's Europa, Saturn's moon Enceladus is covered with an icy crust, under which lies an ocean that could harbour extraterrestrial life. Enceladus is also covered with a unique geological feature: ice volcanoes that spew out water vapour instead of lava!

	Another of Saturn's natural satellites, Titan, is larger than the planet Mercury and is the only moon in the solar system with its own atmosphere.

	In addition to large named moons, anywhere from dozens to hundreds of smaller unnamed bodies called moonlets are embedded inside Saturn's ring system`,
}

var UranusPlanet = Planet{
	Type:                   "Ice giant",
	Size:                   50724,
	RelatedSizeOfEarth:     4,
	Mass:                   "8.7 × 10^25",
	RealtedMassOfEarth:     15.0,
	LengthOfAYear:          84,
	NumberOfMoons:          27,
	AverageDistanceFromSun: 2.9,
	Temperature:            "-226~-197",
	Formation:              `Uranus is the seventh planet from the Sun and is considered an ice giant because it is largely made up of water, ammonia, and methane in solid form. Like all the other planets, Uranus was likely formed in a massive, ancient cloud of gas, dust, and ice that collapsed into a spinning disc. Our Sun was born at its centre, and the planets were created about 4.5 billion years ago from particles sticking together along rings in the disc. Scientists believe the differences between ice giants and gas giants (like Saturn and Jupiter) are due to their slightly different formation histories.`,
	OribitAndRotation: `A day on Uranus lasts about 17 hours. Interestingly, different parts of Uranus rotate at different speeds: while the interior of the planet spins in that time, its upper atmosphere completes a full rotation three hours faster, which creates some of the most violent winds in the solar system. Near Uranus's poles, they blow in the same direction as the rotation of the planet. Closer to the equator, the winds blow in the opposite direction!

	Uranus takes 84 years to complete one full orbit around the Sun.

	Instead of rotating vertically like all of the other planets, Uranus spins around a horizontal axis. Scientists believe that an Earth-sized planet may have collided with Uranus and knocked it on its side.

	Uranus's dramatic 98-degree tilt gives rise to the most extreme seasons in the solar system. For nearly 21 years, a quarter of the Uranian year, one pole is bathed in sunlight while the other exists in complete darkness. For the next 21 years, both poles receive similar amounts of sunlight as they gradually switch places and the cycle repeats.`,
	Seasons:   "",
	Surface:   `Like the other giant planets in our solar system, Uranus does not have a solid surface. Scientists believe that the interior is made up of a solid rocky core covered by a dense liquid layer of water and ammonia. The atmosphere surrounding the interior of the planet is mostly hydrogen and helium. Despite having the second fastest winds in the solar system (after Neptune), Uranus's atmosphere looks quite calm. Very few clouds are observed in its light blue atmosphere.`,
	Exosphere: "",
	Else:      "Rings - Although Saturn is best known for its elaborate ring system, all of the giant planets in our solar system have rings of their own. Uranus has 13 distinct rings made up of very dark particles, quite different from Saturn's rings, which are mostly water ice. Scientists believe that Uranus's rings formed after the planet took shape, from the leftovers of destroyed moons. The rings are tilted about 90 degrees and match Uranus's rotation.",
}

var NeptunePlanet = Planet{
	Type:                   "Ice giant",
	Size:                   49244,
	RelatedSizeOfEarth:     4.0,
	Mass:                   "1.0 × 10^26",
	RealtedMassOfEarth:     17.0,
	LengthOfAYear:          165,
	NumberOfMoons:          14,
	AverageDistanceFromSun: 4.5,
	Temperature:            "-218~-200",
	Formation:              `The eighth planet from the Sun, Neptune is considered an ice giant because it is largely made up of water, ammonia, and methane in solid form. As our solar system took shape about 4.5 billion years ago, Neptune was likely formed in a massive, ancient cloud of gas, dust, and ice which collapsed into a spinning disc with our Sun at its centre. Scientists believe the differences between ice giants and gas giants (like Saturn and Jupiter) are due to their slightly different formation histories.`,
	OribitAndRotation: `Based on measurements taken by the Voyager 2 probe, different parts of Neptune may rotate at different speeds since the planet is not a solid body. Neptune's equator seems to rotate once every 18 hours while its polar regions spin once every 12 hours. This difference in rotational speed between the different regions is the largest of any planet and causes the strongest winds in the solar system, as fast as 2100 km/h!

	Neptune takes 165 years to complete one full orbit around the Sun.`,
	Seasons: "",
	Surface: `Like the other giant planets in our solar system, Neptune does not have a solid surface. Scientists believe that the interior of the planet is made up of a solid rocky core covered by a hot and dense liquid layer of water and ammonia. The atmosphere surrounding the interior of the planet is made up of mostly hydrogen and helium.

	Neptune's blue colour is caused by the methane found in its atmosphere, which absorbs red light. Scientists are unsure why Uranus and Neptune are different shades of blue despite their very similar atmospheres. Like that of Jupiter, Neptune's atmosphere contains many storm systems including the Great Dark Spot, which is about as wide as Earth.`,
	Exosphere: "",
	Else:      "",
}

var MoonPlanet = Planet{
	Type:                   "",
	Size:                   3476,
	RelatedSizeOfEarth:     0.25,
	Mass:                   "7.347 × 10^2",
	RealtedMassOfEarth:     0.01,
	LengthOfAYear:          27.3,
	NumberOfMoons:          384400,
	AverageDistanceFromSun: 0,
	Temperature:            "-248~123",
	Formation:              `The Moon is a rocky body that orbits planet Earth. It likely formed when a very large (Mars-sized) asteroid struck Earth 4.5 billion years ago, very early in our planet's history. The debris from the impact clumped together to form our natural satellite, the Moon. This shared formation explains why Earth and the Moon are made up of many of the same elements.`,
	OribitAndRotation:      `The Moon is tidally locked with Earth, which means that its rotation on its axis and its orbit around Earth are perfectly synchronized, causing the same side of the Moon to always be facing Earth.`,
	Seasons:                "",
	Surface:                `The surface of the Moon is littered with craters caused by asteroid impacts that have occurred over billions of years. Because the Moon does not have much of an atmosphere, or any liquid water or vegetation, lunar craters do not erode over time and their appearance does not change. This means that the Moon is like a cosmic time capsule. By studying it, scientists can learn more about the history of the entire solar system.`,
	Exosphere:              "",
	Else:                   "",
}
