1. 创建sc
	val sc = SparkContext.getOrCreate() // 其中两个强制的配置master 和 appname在 shell 启动的时候必须指定
 或 val sc = SparkContext.getOrCreate(conf)
 
2. yarn-client 和 yarn-cluster 在spark2.0之后被--delopy-mode client/cluster 替换了

3. Setting Local Properties
	
	sc.setLocalProperty("spark.scheduler.pool", "myPool") 通过这个配置可以手动控制不同action触发的jobs是否要在不同的thread pool中运行。

4. Distribute JARs to workers

	The jar you specify with SparkContext.addJar will be copied to all the worker nodes.
		sc.addJar("build.sbt")
		
	广播变量:这里有个注意的的地方： 就是广播变量时executor级别，为何不是task级别？因为task是work上的thread级，每个thread是共享内存变量，那么没有必要将广播变量共享到task级，所有task依旧是可以实现共享

5. Running Jobs (runJob methods)
	
	7中不同需求下的runJob方法：
	
		runJob[T, U](rdd: RDD[T], func: (TaskContext, Iterator[T]) => U, partitions: Seq[Int], resultHandler: (Int, U) => Unit): Unit
		
		runJob[T, U](rdd: RDD[T], func: (TaskContext, Iterator[T]) => U, partitions: Seq[Int]): Array[U]
		
		runJob[T, U](rdd: RDD[T], func: Iterator[T] => U, partitions: Seq[Int]): Array[U]
		
		runJob[T, U](rdd: RDD[T], func: (TaskContext, Iterator[T]) => U): Array[U]
		
		runJob[T, U](rdd: RDD[T], func: Iterator[T] => U): Array[U]
		
		
		runJob[T, U](rdd: RDD[T], processPartition: (TaskContext, Iterator[T]) => U, resultHandler: (Int, U) => Unit): Unit
		
		runJob[T, U: ClassTag](rdd: RDD[T], processPartition: Iterator[T] => U, resultHandler: (Int, U) => Unit): Unit
		
6. Transformations： RDDs 是不可变的，不会被修改，transform操作只是将rdd转换成另一个形式的rdd

	1) There are transformations that may trigger jobs, e.g. sortBy, zipWithIndex, etc.
	
	2) RDD lineage graph using toDebugString method.
	
7. Action: AsyncRDDActions,(thanks to the implicit conversion rddToAsyncRDDActions in RDD class). The methods return a FutureAction.
	
	countAsync
	collectAsync
	takeAsync
	foreachAsync
	foreachPartitionAsync
	
	