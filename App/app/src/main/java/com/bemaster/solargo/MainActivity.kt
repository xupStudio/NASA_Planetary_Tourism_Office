package com.bemaster.solargo

import android.os.Bundle
import androidx.fragment.app.FragmentActivity
import com.bemaster.solargo.databinding.ActivityMainBinding
import com.bemaster.solargo.databinding.BottomSheetEarthBinding
import com.bemaster.solargo.databinding.BottomSheetMarsBinding
import com.bemaster.solargo.databinding.BottomSheetMoonBinding
import com.google.android.material.bottomsheet.BottomSheetDialog

class MainActivity : FragmentActivity() {

    private lateinit var binding: ActivityMainBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityMainBinding.inflate(layoutInflater)
        initUI()
        setContentView(binding.root)
    }

    private fun initUI() {
        binding.ivEarth.setOnClickListener {
            BottomSheetDialog(this).apply {
                setContentView(BottomSheetEarthBinding.inflate(layoutInflater).root)
            }.show()
        }
        binding.ivMoon.setOnClickListener {
            BottomSheetDialog(this).apply {
                setContentView(BottomSheetMoonBinding.inflate(layoutInflater).root)
            }.show()
        }
        binding.ivMars.setOnClickListener {
            BottomSheetDialog(this).apply {
                setContentView(BottomSheetMarsBinding.inflate(layoutInflater).root)
            }.show()
        }
    }
}