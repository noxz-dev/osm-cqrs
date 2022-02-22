FROM ubuntu:20.04 AS compiler-common

FROM compiler-common AS compiler-postgis
RUN apt-get install -y --no-install-recommends \
 postgresql-server-dev-12 \
 libxml2-dev \
 libgeos-dev \
 libproj-dev
RUN wget https://download.osgeo.org/postgis/source/postgis-3.1.1.tar.gz -O postgis.tar.gz \
&& mkdir -p postgis_src \
&& tar -xvzf postgis.tar.gz --strip 1 -C postgis_src \
&& rm postgis.tar.gz \
&& cd postgis_src \
&& ./configure --without-protobuf --without-raster \
&& make -j $(nproc) \
&& checkinstall --pkgversion="3.1.1" --install=no --default make install


RUN wget -https://github.com/omniscale/imposm3/releases/download/v0.11.1/imposm-0.11.1-linux-x86-64.tar.gz -O imposm3.tar.gz \
&& tar -xvzf imposm3.tar.gz \
&& rm imposm3.tar.gz \

RUN adduser --disabled-password --gecos "" renderer

# Configure PosgtreSQL
COPY postgresql.custom.conf /etc/postgresql/12/main/
RUN chown -R postgres:postgres /var/lib/postgresql \
&& chown postgres:postgres /etc/postgresql/12/main/postgresql.custom.conf \
&& echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/12/main/pg_hba.conf \
&& echo "host all all ::/0 md5" >> /etc/postgresql/12/main/pg_hba.conf


COPY run.imposm3.sh /
ENTRYPOINT ["run.imposm3.sh"]
CMD []
EXPOSE 5432