day_1:
	make -C day_1 run

day_2:
	make -C day_2 run

day_3:
	make -C day_3 run

all: day_1 day_2 day_3

.PHONY: all day_1 day_2 day_3