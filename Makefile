day_1:
	make -C day_1 run

day_2:
	make -C day_2 run

all: day_1 day_2

.PHONY: all day_1 day_2