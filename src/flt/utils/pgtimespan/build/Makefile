EXTENSION = pgtimespan
DATA = pgtimespan--0.1.sql  # script files to install
# REGRESS = pgtimespan_test     # our test script file (without extension)
MODULES = pgtimespan          # our c module file to build

# postgres build stuff
PG_CONFIG = pg_config
PGXS := $(shell $(PG_CONFIG) --pgxs)
include $(PGXS)